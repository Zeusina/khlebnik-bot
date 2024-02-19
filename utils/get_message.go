package utils

import (
	"encoding/json"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

func GetMessage(id string) (msg string) {
	workdir, err := os.Getwd()
	if err != nil {
		log.Error(err)
		return "Произошла непредвиденная ошибка, попробуйте еще раз."
	}
	path := filepath.Join(workdir, "assets", "messages.json")
	log.Info(path)
	messagesFile, err := os.ReadFile(path)
	if err != nil {
		log.Error(err)
		return "Произошла непредвиденная ошибка, попробуйте еще раз."
	}
	if err != nil {
		log.Error(err)
		return "Произошла непредвиденная ошибка, попробуйте еще раз."
	}
	var messagesJson MessagesFile
	err = json.Unmarshal(messagesFile, &messagesJson)
	if err != nil {
		log.Error(err)
		return "Произошла непредвиденная ошибка, попробуйте еще раз."
	}
	for _, messageEntry := range messagesJson.Messages.Ru {
		if messageEntry.ID == id {
			return messageEntry.Message
		}
	}
	return "Произошла непредвиденная ошибка, попробуйте еще раз."
}

type MessagesFile struct {
	Messages messages `json:"messages"`
}
type messages struct {
	Ru []messageType `json:"ru"`
}

type messageType struct {
	ID      string `json:"id"`
	Message string `json:"message"`
}
