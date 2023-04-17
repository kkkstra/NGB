package logrus

import (
	"NGB/internal/config"
	"fmt"
	"path"
	"path/filepath"
	"runtime"
	"time"

	nested "github.com/antonfisher/nested-logrus-formatter"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func init() {
	getLogger()
}

func getLogger() {
	Logger = logrus.New()
	setLogger()
}

func setLogger() {
	// 设置格式
	Logger.SetReportCaller(true)
	if config.C.App.Debug {
		Logger.SetLevel(logrus.DebugLevel)
	} else {
		Logger.SetLevel(logrus.InfoLevel)
	}
	Logger.SetFormatter(customFormatter(false))

	// 日志分割
	logPath := path.Join(config.C.Log.Filepath, config.C.Log.FilenamePrefix)
	writer, _ := rotatelogs.New(
		logPath+"-%Y-%m-%d",
		rotatelogs.WithLinkName(logPath),
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	// 绑定钩子，输出到文件
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer,
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, customFormatter(true))

	Logger.AddHook(lfHook)
}

func customFormatter(writeToFile bool) *nested.Formatter {
	return &nested.Formatter{
		HideKeys:        true,
		TimestampFormat: "2006-01-02 15:04:05",
		CallerFirst:     false,
		NoColors:        writeToFile,
		CustomCallerFormatter: func(frame *runtime.Frame) string {
			funcInfo := runtime.FuncForPC(frame.PC)
			if funcInfo == nil {
				return "error during runtime.FuncForPC"
			}
			fullPath, line := funcInfo.FileLine(frame.PC)
			return fmt.Sprintf(" [%v:%v]", filepath.Base(fullPath), line)
		},
	}
}
