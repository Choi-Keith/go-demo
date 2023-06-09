package logrusx

import (
	"context"
	"os"
	"runtime"
	"strconv"

	"github.com/sirupsen/logrus"
)

func InitLog(flag string) {
	if flag == "production" {
		logrus.SetFormatter(&logrus.JSONFormatter{})
		fs, err2 := os.Create("./log")
		if err2 != nil {
			return
		}
		logrus.SetOutput(fs)
	} else {
		logrus.SetFormatter(&logrus.TextFormatter{})
		logrus.SetOutput(os.Stdout)
	}
	level, err := logrus.ParseLevel("info")
	if err != nil {
		panic(err)
	}
	logrus.SetLevel(level)
}

const (
	traceKey = "trace"
	fileKey  = "File"
)

func getCallerInfo() (fileline string) {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		panic("Could not get context info for logger")
	}
	filename := file + ":" + strconv.Itoa(line)
	return filename
}

func Info(ctx context.Context, args ...interface{}) {
	val, ok := ctx.Value(traceKey).(string)
	if !ok {
		val = ""
	}
	filename := getCallerInfo()
	logrus.WithField(traceKey, val).WithField(fileKey, filename).Info(args...)
}

func Infof(ctx context.Context, format string, args ...interface{}) {
	val, ok := ctx.Value(traceKey).(string)
	if !ok {
		val = ""
	}
	filename := getCallerInfo()
	logrus.WithField(traceKey, val).WithField(fileKey, filename).Infof(format, args...)
}

func Debug(ctx context.Context, args ...interface{}) {
	val, ok := ctx.Value(traceKey).(string)
	if !ok {
		val = ""
	}
	filename := getCallerInfo()
	logrus.WithField(traceKey, val).WithField(fileKey, filename).Debug(args...)
}

func Debugf(ctx context.Context, format string, args ...interface{}) {
	val, ok := ctx.Value(traceKey).(string)
	if !ok {
		val = ""
	}
	filename := getCallerInfo()
	logrus.WithField(traceKey, val).WithField(fileKey, filename).Debugf(format, args...)
}

func Warming(ctx context.Context, args ...interface{}) {
	val, ok := ctx.Value(traceKey).(string)
	if !ok {
		val = ""
	}
	filename := getCallerInfo()
	logrus.WithField(traceKey, val).WithField(fileKey, filename).Warning(args...)
}

func Warmingf(ctx context.Context, format string, args ...interface{}) {
	val, ok := ctx.Value(traceKey).(string)
	if !ok {
		val = ""
	}
	filename := getCallerInfo()
	logrus.WithField(traceKey, val).WithField(fileKey, filename).Warningf(format, args...)
}

func Error(ctx context.Context, args ...interface{}) {
	val, ok := ctx.Value(traceKey).(string)
	if !ok {
		val = ""
	}
	filename := getCallerInfo()
	logrus.WithField(traceKey, val).WithField(fileKey, filename).Error(args...)
}

func Errorf(ctx context.Context, format string, args ...interface{}) {
	val, ok := ctx.Value(traceKey).(string)
	if !ok {
		val = ""
	}
	filename := getCallerInfo()
	logrus.WithField(traceKey, val).WithField(fileKey, filename).Errorf(format, args...)
}
