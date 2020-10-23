package calendar

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

var (
	baseEndpoint = "https://www.googleapis.com/calendar/v3/calendars/"
)

type Events struct {
	Summary string `json:"summary"`
	Items   []Item `json:"items"`
}

type Item struct {
	Summary  string       `json:"summary"`
	HtmlLink string       `json:"htmlLink"`
	Start    calendarTime `json:"start"`
	End      calendarTime `json:"end"`
}

type calendarTime struct {
	DateTime string `json:"dateTime"`
	TimeZone string `json:"timeZone"`
}

func (calTime calendarTime) String() string {
	return calTime.DateTime
}

func constructCalendarEventsEndpoint(calendarRequest *Request) (calendarEventsEndpoint string) {
	baseUrl, err := url.Parse(baseEndpoint + calendarRequest.Id + "/events")
	if err != nil {
		fmt.Println("Malformed URL: ", err.Error())
		return
	}

	timeMin := calendarRequest.TimeMin.Format(time.RFC3339)
	timeMax := calendarRequest.TimeMax.Format(time.RFC3339)

	params := url.Values{}
	params.Add("orderBy", "startTime")
	params.Add("singleEvents", "true")
	params.Add("timeMin", timeMin)
	params.Add("timeMax", timeMax)
	params.Add("timeZone", TimeZone)
	params.Add("key", Token)

	baseUrl.RawQuery = params.Encode()

	return baseUrl.String()
}

func RequestCalendarEvents(calendarRequest *Request) (events *Events) {
	calendarUrl := constructCalendarEventsEndpoint(calendarRequest)

	resp, err := http.Get(calendarUrl)
	if err != nil {
		log.Printf("error while fetching calendar: %s", err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("error while fetching calendar's body: %s", err)
		return
	}

	err = json.Unmarshal(body, &events)
	if err != nil {
		log.Printf("error while filling up events struct: %s", err)
		return
	}

	return
}

func GenerateTextResponse(events *Events, inputDate string) (textResponse string) {
	for _, item := range events.Items {
		if len(item.Summary) == 0 {
			continue
		}

		var (
			startTime time.Time
			endTime   time.Time
		)

		err := json.Unmarshal([]byte(fmt.Sprintf("\"%s\"", item.Start.String())), &startTime)
		if err != nil {
			log.Printf("start time unmarshall error: %s", err)
		}
		err = json.Unmarshal([]byte(fmt.Sprintf("\"%s\"", item.End.String())), &endTime)
		if err != nil {
			log.Printf("end time unmarshall error: %s", err)
		}

		startTimeParsed := startTime.Format("15:04")
		endTimeParsed := endTime.Format("15:04")

		durationTime := endTime.Sub(startTime)
		durationParsed := fmt.Sprintf("%dh", int(durationTime.Hours()))

		textResponse += fmt.Sprintf(
			"\n\n\t\t<b>\"%s\"</b>\n\t\t<code>%s godz. %s do %s (%s)</code>\n\t\t<a href=\"%s\">Sprawdź w kalendarzu</a>",
			item.Summary,
			inputDate[:len(inputDate)-5],
			startTimeParsed,
			endTimeParsed,
			durationParsed,
			item.HtmlLink,
		)
	}

	if len(textResponse) > 0 {
		textPrefix := fmt.Sprintf(
			"<b>Wyniki wyszukiwania dla kalendarza \"%s\" na dzień %s</b>\n(Kliknij wybraną godzinę aby skopiować)",
			events.Summary,
			inputDate,
		)
		textResponse = textPrefix + textResponse
	} else {
		textResponse = fmt.Sprintf(
			"<b>Brak wolnych terminów w kalendarzu \"%s\" na dzień %s</b>",
			events.Summary,
			inputDate,
		)
	}

	return
}
