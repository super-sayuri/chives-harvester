package conf

import (
	"context"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

type LogConfig struct {
	Format string `yaml:"format"`
	Output string `yaml:"output"`
	Level  string `yaml:"level"`
	Path   string `yaml:"path"`
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
	if strings.ToUpper(conf.Output) == "FILE" {
		if len(conf.Path) == 0 {
			conf.Path = "syr.log"
		}
		f, err := os.OpenFile(conf.Path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
		if err != nil {
			return err
		}
		logrus.SetOutput(f)
	} else {
		logrus.SetOutput(os.Stdout)
	}
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
