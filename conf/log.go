package conf

import (
	"context"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

type LogConfig struct {
	Format string `yml:"format"`
	Output string `yml:"output"`
	Level  string `yml:"level"`
}

const (
	LOG_KEY_JOB_ID     = "bId"
	LOG_KEY_REQUEST_ID = "requestId"
)

func logInit(conf *LogConfig) error {
	if conf == nil {
		conf = &LogConfig{
			Level: "INFO",
		}
	}
	if strings.ToUpper(conf.Format) == "JSON" {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	} else {
		logrus.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true,
		})
	}
	// todo add file output
	logrus.SetOutput(os.Stdout)
	level, err := logrus.ParseLevel(conf.Level)
	if err != nil {
		return err
	}
	logrus.SetLevel(level)
	logrus.SetReportCaller(true)
	return nil
}

func GetLog(ctx context.Context) *logrus.Entry {
	var l *logrus.Entry
	if ctx.Value(LOG_KEY_JOB_ID) != nil {
		l = logrus.WithField(LOG_KEY_JOB_ID, ctx.Value(LOG_KEY_JOB_ID))
	}
	if ctx.Value(LOG_KEY_REQUEST_ID) != nil {
		l = logrus.WithField(LOG_KEY_REQUEST_ID, ctx.Value(LOG_KEY_REQUEST_ID))
	}
	if l == nil {
		l = logrus.NewEntry(logrus.StandardLogger())
	}
	return l
}
