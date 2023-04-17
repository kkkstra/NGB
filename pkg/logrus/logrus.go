package logrus

import (
	"fmt"
	"path/filepath"
	"runtime"
	"time"

	nested "github.com/antonfisher/nested-logrus-formatter"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func InitLogger(debug bool, logPath string) {
	getLogger(debug, logPath)
}

func getLogger(debug bool, logPath string) {
	Logger = logrus.New()
	setLogger(debug, logPath)
}

func setLogger(debug bool, logPath string) {
	// 设置格式
	Logger.SetReportCaller(true)
	if debug {
		Logger.SetLevel(logrus.DebugLevel)
	} else {
		Logger.SetLevel(logrus.InfoLevel)
	}
	Logger.SetFormatter(customFormatter(false))

	// 日志分割
	// logPath := path.Join(filepath, filenamePrefix)
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
