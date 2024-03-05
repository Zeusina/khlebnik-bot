package handlers

import (
	"context"

	"github.com/Zeusina/khlebnik-bot/utils"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	log "github.com/sirupsen/logrus"
)

func QuesionRequestHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	log.Debugf("User %d wants to ask question", update.Message.From.ID)
	utils.SetFSM("askquestion", update.Message.From.ID)
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   utils.GetMessage("questionrequest"),
		ReplyMarkup: models.ReplyKeyboardRemove{
			RemoveKeyboard: true,
		},
	})
}
