package initialize

import (
	"admin-cli/config"
	"admin-cli/global"
	"context"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"io"
	"log"
	"os"
	"time"
)

// InitConfig 初始化config
func InitConfig() error {
	viper.AddConfigPath("./config/")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Error("配置文件读取err", err)
		return err
	}
	var configData config.Config
	if err = viper.Unmarshal(&configData); err != nil {
		logrus.Error("配置文件 Unmarshal err", err)
		return err
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		logrus.Info("配置文件发生更新")
		if err := viper.Unmarshal(&configData); err != nil {
			logrus.Infof("配置文件更新解析失败,err:%v", err)
		}
	})
	config.SetConfig(&configData)
	return nil
}

// InitLog 初始化日志配置
func InitLog() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetReportCaller(true)
	logrus.SetLevel(logrus.InfoLevel)
	cfg := config.GetConfig()
	path := "./" + cfg.Log.Path + "/log.log"
	logger := &lumberjack.Logger{
		LocalTime:  true,
		Filename:   path,
		MaxSize:    cfg.Log.Maxsize, // megabytes
		MaxBackups: cfg.Log.MaxBackups,
		MaxAge:     cfg.Log.MaxAge,   //days
		Compress:   cfg.Log.Compress, // disabled by default
	}
	writers := []io.Writer{
		logger,
		os.Stdout,
	}
	fileAndStdoutWriter := io.MultiWriter(writers...)
	logrus.SetOutput(fileAndStdoutWriter)
}

// InitMysql 初始化数据库
func InitMysql() error {

	configData := config.GetConfig()

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local",
		configData.Mysql.UserName,
		configData.Mysql.Password,
		configData.Mysql.Host,
		configData.Mysql.Port,
		configData.Mysql.DataBase,
	)
	path := "./" + configData.Log.Path + "/mysql.log"
	sqlLogger := &lumberjack.Logger{
		LocalTime:  true,
		Filename:   path,
		MaxSize:    configData.Log.Maxsize, // megabytes
		MaxBackups: configData.Log.MaxBackups,
		MaxAge:     configData.Log.MaxAge,   //days
		Compress:   configData.Log.Compress, // disabled by default
	}
	writers := []io.Writer{
		sqlLogger,
		os.Stdout,
	}
	fileAndStdoutWriter := io.MultiWriter(writers...)
	newLogger := logger.New(
		log.New(fileAndStdoutWriter, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,       // 禁用彩色打印
		},
	)

	Db, err := gorm.Open(mysql.Open(dataSource), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})

	if err != nil {
		return err
	}
	sqlDB, err := Db.DB()
	if err != nil {
		logrus.Error(err)
		return err
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(50)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(1 * time.Hour)

	err = Db.AutoMigrate()
	if err != nil {
		logrus.Errorf("MYSQL AutoMigrate error %v", err)
		return err
	}
	global.Db = Db
	return nil
}

//初始化redis
func InitRedis() error {
	configData := config.GetConfig()
	redisConfig := &redis.Options{
		Addr:     fmt.Sprintf("%s:%s", configData.Redis.Host, configData.Redis.Port),
		Password: configData.Redis.Password,
		DB:       configData.Redis.DB,
	}
	client := redis.NewClient(redisConfig)
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		logrus.Errorf("redis connect error %v", err)
		return err
	}
	global.Redis = client
	return nil
}
