package models

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type Message struct {
	Text string `json:"text"`
	Chat Chat   `json:"chat"`
}

type SendMessageRequest struct {
	Text string `json:"text"`
	ChatId string `json:"chat_id"`
}

func SendMessage(message *Message) (err error) {
	request := &SendMessageRequest{
		Text:   message.Text,
		ChatId: message.Chat.Id,
	}
	reqBytes, err := json.Marshal(request)
	if err != nil {
		return
	}

	response, err := http.Post(TelegramSendMessageEndpoint, "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return
	}

	if response.StatusCode != http.StatusOK {
		err = errors.New("unexpected status" + response.Status)
	}

	log.Printf("message sent successfully")

	return
}
