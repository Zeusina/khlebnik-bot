package handlers

import (
	"context"

	"github.com/Zeusina/khlebnik-bot/utils"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	log "github.com/sirupsen/logrus"
)

func QuestionFSMCheck(update *models.Update) bool {
	log.Debugf("Current FSM state - %s", utils.GetFSM(update.Message.From.ID))
	return utils.GetFSM(update.Message.From.ID) == "askquestion"
}

func QuestionHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	log.Debugf("%d used /start command", update.Message.From.ID)
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   utils.GetMessage("questionaccepted"),
	})
	utils.SetFSM("none", update.Message.From.ID)
	// TODO: add interaction with model here
}
