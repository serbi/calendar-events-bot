package models

import "os"

var (
	token = os.Getenv("TOKEN")

	TelegramBotApiUrl = "https://api.telegram.org/bot"
	TelegramSendMessageEndpoint = TelegramBotApiUrl + token + "/sendMessage"
)
