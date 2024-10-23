package logger

import (
	"os"

	log "github.com/sirupsen/logrus"
)

var e *log.Entry

type Logger struct {
	*log.Entry
}

func GetLogger() Logger {
	return Logger{e}
}

func LogInit() *log.Logger {
	l := log.New()
	l.SetReportCaller(true)
	l.Formatter = &log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	}

	l.SetOutput(os.Stdout)

	e = log.NewEntry(l)
	return l
}
