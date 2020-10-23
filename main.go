package main

import (
	"log"
	"net/http"

	"github.com/serbi/calendar_events_bot/routes"
)

func main() {
	//var port = os.Getenv("PORT")
	var addr = "localhost:" + "3000"

	http.HandleFunc("/webhook", routes.WebhookHandler)

	log.Printf("Listening on %s ...", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
