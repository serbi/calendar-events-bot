package webhook

import "github.com/serbi/calendar_events_bot/models/telegram"

type Request struct {
	telegram.Message `json:"message"`
}
