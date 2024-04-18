package handlers

import (
	"context"

	"github.com/Zeusina/khlebnik-bot/utils"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	log "github.com/sirupsen/logrus"
)

func StartHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	log.WithField("userid", update.Message.From.ID).Debug("Bot started")
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   utils.GetMessage("startmessage"),
		ReplyMarkup: models.ReplyKeyboardMarkup{
			Keyboard: [][]models.KeyboardButton{
				{
					{
						Text: utils.GetMessage("askquestion"),
					},
				},
			},
			ResizeKeyboard: true,
		},
		ParseMode: models.ParseModeHTML,
	})
	if err != nil {
		log.Error(err)
	}
}
