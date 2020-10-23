package actions

import (
	"log"
	"regexp"
	"time"

	"github.com/serbi/calendar_events_bot/models/calendar"
)

func CalendarEventsAction(message string) (response string) {
	dateReg := regexp.MustCompile(`([0-9]{2}|[0-9])-[0-9]{2}-[0-9]{4}`)
	dateInput := dateReg.FindString(message)

	if dateInput == "" {
		return
	}

	timeMin, err := time.Parse("02-01-2006", dateInput)
	if err != nil {
		log.Printf("error while parsing timeMin value: %s", err)
		return
	}
	timeMax := timeMin.AddDate(0, 0, 1)

	calendarReq := &calendar.Request{
		Id:      calendar.SalaSofaId,
		TimeMin: timeMin,
		TimeMax: timeMax,
	}

	events := calendar.RequestCalendarEvents(calendarReq)

	return calendar.GenerateTextResponse(events, dateInput)
}
