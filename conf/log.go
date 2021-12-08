package conf

import (
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

type LogConfig struct {
	Format string `yml:"format"`
	Output string `yml:"output"`
	Level  string `yml:"level"`
}

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
