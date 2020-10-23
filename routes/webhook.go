package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/serbi/calendar_events_bot/models"
	"github.com/serbi/calendar_events_bot/models/telegram"
	"github.com/serbi/calendar_events_bot/models/telegram/webhook"
)

func WebhookHandler(res http.ResponseWriter, req *http.Request) {
	request := &webhook.Request{}
	if err := json.NewDecoder(req.Body).Decode(request); err != nil {
		log.Printf("error in decoding request body: %s", err)
		return
	}

	responseMessage := models.CompareTextMessageAgainstCommands(request.Message.Text)

	if responseMessage == "" {
		return
	}

	message := &telegram.Message{
		Text: responseMessage,
		Chat: telegram.Chat{
			Id: request.Message.Chat.Id,
		},
	}

	if err := telegram.SendMessage(message); err != nil {
		log.Printf("error in sending message: %s", err)
		return
	}
}
