package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	_ "github.com/lib/pq"
	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

var (
	bot      *tgbotapi.BotAPI
	port     = os.Getenv("PORT")
	botToken = os.Getenv("TOKEN")
	baseURL  = os.Getenv("PUBLIC_URL")
)

func initTelegram() {
	var err error

	bot, err = tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Println(err)
		return
	}

	// this perhaps should be conditional on GetWebhookInfo()
	// only set webhook if it is not set properly
	url := baseURL + bot.Token
	_, err = bot.SetWebhook(tgbotapi.NewWebhook(url))
	if err != nil {
		log.Println(err)
	}
}

func webhookHandler(c *gin.Context) {
	defer c.Request.Body.Close()

	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Println(err)
		return
	}

	var update tgbotapi.Update
	err = json.Unmarshal(bytes, &update)
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("From: %+v Text: %+v\n", update.Message.From, update.Message.Text)
}

func main() {
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())

	initTelegram()
	router.POST("/"+bot.Token, webhookHandler)

	err := router.Run(":" + port)
	if err != nil {
		log.Println(err)
	}
}
