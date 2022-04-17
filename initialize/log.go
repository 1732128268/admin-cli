package initialize

import (
	"admin-cli/config"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
)

// InitLog 初始化日志配置
func InitLog() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetReportCaller(true)
	logrus.SetLevel(logrus.InfoLevel)
	cfg := config.GetConfig()
	path := "./" + cfg.Log.Path + "/log.log"
	l := &lumberjack.Logger{
		LocalTime:  true,
		Filename:   path,
		MaxSize:    cfg.Log.Maxsize, // megabytes
		MaxBackups: cfg.Log.MaxBackups,
		MaxAge:     cfg.Log.MaxAge,   //days
		Compress:   cfg.Log.Compress, // disabled by default
	}
	writers := []io.Writer{
		l,
		os.Stdout,
	}
	fileAndStdoutWriter := io.MultiWriter(writers...)
	logrus.SetOutput(fileAndStdoutWriter)
}
