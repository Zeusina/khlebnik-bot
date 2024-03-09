package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

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
	questionMessage, _ := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   utils.GetMessage("questionaccepted"),
	})
	utils.SetFSM("none", update.Message.From.ID)

	host := os.Getenv("HOST")
	if host == "" {
		host = "localhost"
		log.Warn("API host not set, using 'localhost'")
	}
	baseurl, err := url.Parse(fmt.Sprintf("http://%s:8080", host))
	if err != nil {
		log.Error(err)
	}
	baseurl.Path += "predict"
	params := url.Values{}
	params.Add("question", update.Message.Text)
	baseurl.RawQuery = params.Encode()

	question_predict, err := http.Get(baseurl.String())

	if err != nil {
		log.Error(err)
	}
	defer question_predict.Body.Close()
	body, err := io.ReadAll(question_predict.Body)
	if err != nil {
		log.Error(err)
	}

	var bodyStruct []answerPredict

	json.Unmarshal(body, &bodyStruct)
	log.WithField("content", bodyStruct).Debug("get prediction for question")
	if bodyStruct[0].Score >= 0.67 {
		baseurl, err = url.Parse(fmt.Sprintf("http://%s:8080/answer/%d", host, bodyStruct[0].CorpusId))
		if err != nil {
			log.Error(err)
		}

		var answer questionAnswer

		questionAnswerResponse, err := http.Get(baseurl.String())
		if err != nil {
			log.Error(err)
		}
		body, err = io.ReadAll(questionAnswerResponse.Body)
		if err != nil {
			log.Error(err)
		}
		defer questionAnswerResponse.Body.Close()
		err = json.Unmarshal(body, &answer)
		if err != nil {
			log.Error(err)
		}
		log.WithField("content", string(body)).Debug("Get answer for question")

		b.EditMessageText(ctx, &bot.EditMessageTextParams{
			MessageID: questionMessage.ID,
			ChatID:    questionMessage.Chat.ID,
			Text:      answer.Answer,
		})
		log.WithField("messagetext", answer.Answer).Debug("Message edited")
	} else {
		b.EditMessageText(ctx, &bot.EditMessageTextParams{
			MessageID: questionMessage.ID,
			ChatID:    questionMessage.Chat.ID,
			Text:      utils.GetMessage("calloperator"),
		})
		log.WithField("messagetext", utils.GetMessage("calloperator")).Debug("Message edited")
	}

}

type answerPredict struct {
	CorpusId       int     `json:"corpus_id"`
	Score          float64 `json:"score"`
	CorpusQuestion string  `json:"corpus_quesion"`
}

type questionAnswer struct {
	Id       int    `json:"id"`
	Question string `json:"int"`
	Answer   string `json:"answer"`
}
