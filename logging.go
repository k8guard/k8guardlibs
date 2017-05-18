package k8guardlibs

import (
	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func SetLogger() {
	Log = logrus.New()
	switch l := Cfg.LogLevel; l {
	case "info":
		Log.Level = logrus.InfoLevel
	case "debug":
		Log.Level = logrus.DebugLevel
	case "fatal":
		Log.Level = logrus.FatalLevel
	case "error":
		Log.Level = logrus.ErrorLevel
	default:
		Log.Level = logrus.DebugLevel
	}

}
