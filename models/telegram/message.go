package telegram

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

var (
	BotApiUrl           = "https://api.telegram.org/bot"
	SendMessageEndpoint = BotApiUrl + Token + "/sendMessage"
	ParseMode           = "HTML"
)

type Message struct {
	Text string `json:"text"`
	Chat Chat   `json:"chat"`
}

type SendMessageRequest struct {
	Text      string `json:"text"`
	ChatId    int    `json:"chat_id"`
	ParseMode string `json:"parse_mode"`
}

func SendMessage(message *Message) (err error) {
	request := &SendMessageRequest{
		Text:      message.Text,
		ChatId:    message.Chat.Id,
		ParseMode: ParseMode,
	}
	reqBytes, err := json.Marshal(request)
	if err != nil {
		return
	}

	response, err := http.Post(SendMessageEndpoint, "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return
	}

	if response.StatusCode != http.StatusOK {
		err = errors.New("unexpected status: " + response.Status)
		return
	}

	log.Printf("message sent successfully")

	return
}
