package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/serbi/calendar_events_bot/models"
)

func WebhookHandler(res http.ResponseWriter, req *http.Request) {
	request := &models.WebhookRequest{}
	if err := json.NewDecoder(req.Body).Decode(request); err != nil {
		log.Printf("error in decoding request body: %s", err)
		return
	}

	if !strings.Contains(strings.ToLower(request.Message.Text), PingTextMessage) {
		return
	}

	message := &models.Message{
		Text: PongTextMessage,
		Chat: models.Chat{
			Id: request.Message.Chat.Id,
		},
	}

	if err := models.SendMessage(message); err != nil {
		log.Printf("error in sending message: %s", err)
		return
	}

	log.Printf("sent message")
}
