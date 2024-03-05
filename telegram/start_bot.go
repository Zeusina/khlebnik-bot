package telegram

import (
	"context"
	"os"
	"os/signal"

	"github.com/Zeusina/khlebnik-bot/telegram/handlers"
	"github.com/Zeusina/khlebnik-bot/utils"
	"github.com/go-telegram/bot"
	log "github.com/sirupsen/logrus"
)

func StartBot() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	var opts []bot.Option
	if os.Getenv("DEBUG") == "1" {
		opts = append(opts, bot.WithDebug())
		log.Debug("Configure bot to run in debug mode")
	}

	b, err := bot.New(os.Getenv("TOKEN"), opts...)
	if err != nil {
		panic(err)
	}
	utils.InitFSM()
	b.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypePrefix, handlers.StartHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, utils.GetMessage("askquestion"), bot.MatchTypeExact, handlers.QuesionRequestHandler)
	b.RegisterHandlerMatchFunc(handlers.QuestionFSMCheck, handlers.QuestionHandler)

	b.Start(ctx)
}
