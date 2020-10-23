package models

import (
	"strings"

	"github.com/serbi/calendar_events_bot/actions"
)

type CommandAction func(string) string

type Command struct {
	TextMessage string
	Action      CommandAction
}

func CompareTextMessageAgainstCommands(message string) (messageResponse string) {
	var commands = []Command{
		{
			TextMessage: PingTextCommandMessage,
			Action: func(message string) string {
				return actions.PingAction()
			},
		},
		{
			TextMessage: CalendarEventsCommandMessage,
			Action: func(message string) string {
				return actions.CalendarEventsAction(message)
			},
		},
	}

	for _, command := range commands {
		if strings.Contains(message, command.TextMessage) {
			messageResponse = command.Action(message)
			return
		}
	}

	return ""
}
