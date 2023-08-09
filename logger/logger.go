package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

func Init() {
	log := logrus.New()
	log.SetLevel(logrus.WarnLevel)
	file, err := os.OpenFile("log.log", os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		log.Out = file
	} else {
		log.Info("Failed to log to file, using default stderr")
	}
}
