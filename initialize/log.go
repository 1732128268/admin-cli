package initialize

import (
	"admin-cli/global"
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
	path := "./" + global.Config.Log.Path + "/log.log"
	l := &lumberjack.Logger{
		LocalTime:  true,
		Filename:   path,
		MaxSize:    global.Config.Log.Maxsize, // megabytes
		MaxBackups: global.Config.Log.MaxBackups,
		MaxAge:     global.Config.Log.MaxAge,   //days
		Compress:   global.Config.Log.Compress, // disabled by default
	}
	writers := []io.Writer{
		l,
		os.Stdout,
	}
	fileAndStdoutWriter := io.MultiWriter(writers...)
	logrus.SetOutput(fileAndStdoutWriter)
}
