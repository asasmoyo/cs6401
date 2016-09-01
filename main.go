package main

import (
	"log"
	"os"

	"github.com/kataras/iris"
	telegramApi "gopkg.in/telegram-bot-api.v4"

	"github.com/asasmoyo/cs6401/libs"
)

// AppURL value
var appURL = os.Getenv("APP_URL")
var port = os.Getenv("PORT")

func init() {
	if port == "" {
		port = "8080"
	}
	if appURL == "" {
		log.Panic("Cannot read APP_URL from env vars!")
	}
}

func main() {
	libs.SetupBot(appURL)

	iris.Post("/telegramHandler", func(ctx *iris.Context) {
		update := &telegramApi.Update{}
		ctx.ReadJSON(update)
		log.Printf("Got %d: %s", update.UpdateID, update.Message.Text)
	})

	iris.Listen(":" + port)
}
