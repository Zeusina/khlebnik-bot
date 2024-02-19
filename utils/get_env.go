package utils

import (
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func GetEnv() {
	if os.Getenv("DOCKER") == "1" {
		log.Info("Working in Docker, not getting .env")
		return
	}
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Succesfully load .env")
}
