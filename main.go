package main

import (
	"net/http"
	"os"

	"github.com/serbi/calendar_events_bot/routes"
)

var port = os.Getenv("PORT")
var webhookAddr = ":" + port + "/webhook"

func main() {
	err := http.ListenAndServe(webhookAddr, http.HandlerFunc(routes.WebhookHandler))
	if err != nil {
		panic(err)
	}
}
