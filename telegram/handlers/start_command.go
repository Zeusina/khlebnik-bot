package handlers

import (
	"context"

	"github.com/Zeusina/khlebnik-bot/utils"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	log "github.com/sirupsen/logrus"
)

func StartHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	log.Debugf("%d used /start command", update.Message.From.ID)
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   utils.GetMessage("startmessage"),
	})
}
