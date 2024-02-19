package utils

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func ConfigureLogs() {
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		ForceColors:   true,
		FullTimestamp: true,
	})
	log.Debug("Started log configuring")
	var logLevel string = os.Getenv("LOG_LEVEL")
	switch logLevel {
	case "DEBUG":
		log.SetLevel(log.DebugLevel)
	case "ERROR":
		log.SetLevel(log.ErrorLevel)
	case "FATAL":
		log.SetLevel(log.FatalLevel)
	default:
		log.Warn("Log level not set, using ERROR")
		log.SetLevel(log.ErrorLevel)
	}

	log.Info("Succesfully configured logs")
}
