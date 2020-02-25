package logger

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/wendellliu/good-search/pkg/config"
)

var Logger = logrus.New()

func Load() {
	if config.Config.DevelopmentMode {
		Logger.SetLevel(logrus.DebugLevel)
	}
	Logger.SetOutput(os.Stdout)
	Logger.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
}
