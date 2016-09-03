package main

import (
	"log"
	"os"
	"strings"

	"github.com/kataras/iris"
	telegramApi "gopkg.in/telegram-bot-api.v4"

	"github.com/asasmoyo/cs6401/lib/telegram"
)

// AppURL value
var appURL = os.Getenv("APP_URL")
var telegramToken = os.Getenv("TELEGRAM_TOKEN")
var port = os.Getenv("PORT")

func init() {
	if appURL == "" {
		log.Panic("Cannot read APP_URL from env vars!")
	}
	if telegramToken == "" {
		log.Panic("Cannot read TELEGRAM_TOKEN from env vars!")
	}
	if port == "" {
		port = "8080"
	}
}

func main() {
	bot := telegram.NewBot(telegramToken, appURL)

	iris.Post("/telegramHandler", func(ctx *iris.Context) {
		update := &telegramApi.Update{}
		ctx.ReadJSON(update)
		log.Printf("Got %d: %s", update.UpdateID, update.Message.Text)

		chatResponse := handleUpdate(update)
		log.Printf("Sending reply to chatid: %d message: %s", update.Message.Chat.ID, chatResponse)
		bot.SendMessage(update.Message.Chat.ID, chatResponse)
	})

	iris.Listen(":" + port)
}

func handleUpdate(update *telegramApi.Update) string {
	textMessage := update.Message.Text
	var response string

	if strings.HasPrefix(textMessage, "/task1") {
		parts := strings.Split(textMessage, " ")
		if len(parts) != 3 || !(checkCleanOption(parts[1]) && checkPositionOption(parts[2])) {
			response = "Invalid parameters"
		} else {
			response = "Working on it"
		}
	} else {
		response = `Hi there! This is ArbaCS6401Bot. You can ask for tasks specified bellow:
      - /task1 [is clean: YES/NO] [position: LEFT/RIGHT]`
	}

	return response
}

func checkCleanOption(value string) bool {
	if value == "yes" || value == "no" {
		return true
	}
	return false
}

func checkPositionOption(value string) bool {
	if value == "left" || value == "right" {
		return true
	}
	return false
}
