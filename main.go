package main

import (
	"log"
	"net/http"
	"os"

	"github.com/serbi/calendar_events_bot/routes"
)

func main() {
	var port = os.Getenv("PORT")

	http.HandleFunc("/webhook", routes.WebhookHandler)

	log.Printf("Listening on %s ...", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
