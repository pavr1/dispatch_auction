package pkg

import (
	"os"

	"github.com/sirupsen/logrus"
)

func InitializeLogging() *logrus.Logger {
	log := logrus.New()
	log.SetOutput(os.Stdout)
	//log.SetReportCaller(true)
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetLevel(logrus.DebugLevel)

	return log
}
