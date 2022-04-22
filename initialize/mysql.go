package initialize

import (
	"admin-cli/global"
	"admin-cli/model"
	"fmt"
	adapter "github.com/casbin/gorm-adapter/v3"
	"github.com/sirupsen/logrus"
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

// InitMysql 初始化数据库
func InitMysql() error {

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local",
		global.Config.Mysql.UserName,
		global.Config.Mysql.Password,
		global.Config.Mysql.Host,
		global.Config.Mysql.Port,
		global.Config.Mysql.DataBase,
	)
	path := "./" + global.Config.Log.Path + "/mysql.log"
	sqlLogger := &lumberjack.Logger{
		LocalTime:  true,
		Filename:   path,
		MaxSize:    global.Config.Log.Maxsize, // megabytes
		MaxBackups: global.Config.Log.MaxBackups,
		MaxAge:     global.Config.Log.MaxAge,   //days
		Compress:   global.Config.Log.Compress, // disabled by default
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

	err = Db.AutoMigrate(
		&model.User{},
		&model.RoleAuthority{},
		&model.SysBaseMenu{},
		&adapter.CasbinRule{},
		&model.SysOperationRecord{},
	)
	if err != nil {
		logrus.Errorf("MYSQL AutoMigrate error %v", err)
		return err
	}
	global.Db = Db
	return nil
}
