package main

import (
	"net/http"
	"os"

	"github.com/serbi/calendar_events_bot/routes"
)

func main() {
	var port = os.Getenv("PORT")

	err := http.ListenAndServe(":"+port, http.HandlerFunc(routes.WebhookHandler))
	if err != nil {
		panic(err)
	}
}
