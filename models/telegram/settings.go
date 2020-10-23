package telegram

import "os"

var (
	Token = os.Getenv("TELEGRAM_TOKEN")
)
