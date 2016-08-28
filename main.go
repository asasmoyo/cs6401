package main

import (
	"github.com/kataras/iris"
	telegramApi "gopkg.in/telegram-bot-api.v4"
	"log"
	"os"

	"github.com/asasmoyo/cs6401/libs"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	appURL := os.Getenv("APP_URL")
	if appURL == "" {
		log.Panic("Cannot read APP_URL from env vars!")
	}

	telegramToken := os.Getenv("TELEGRAM_TOKEN")
	libs.SetupBot(telegramToken, appURL)

	iris.Post("/telegramHandler", func(ctx *iris.Context) {
		response := &telegramApi.Update{}
		ctx.ReadJSON(response)
		log.Printf("Got %d: %s", response.UpdateID, response.Message.Text)
	})

	iris.Listen(":" + port)
}
