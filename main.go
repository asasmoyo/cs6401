package main

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/kataras/iris"
	telegramApi "gopkg.in/telegram-bot-api.v4"

	"github.com/asasmoyo/cs6401/lib/task1"
	"github.com/asasmoyo/cs6401/lib/telegram"
)

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
		log.Println("No PORT env var detected, bind app on port 8080")
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
		if len(parts) != 4 || !(checkPositionOption(parts[1]) && checkCleanOption(parts[2]) && checkCleanOption(parts[3])) {
			response = "Invalid parameters"
		} else {
			position := strings.ToLower(parts[1])
			isLeftClean, _ := strconv.ParseBool(strings.ToLower(parts[2]))
			isRightClean, _ := strconv.ParseBool(strings.ToLower(parts[3]))
			state := task1.GetStep(position, isLeftClean, isRightClean)
			response = "You are now at state #" + strconv.Itoa(state.No) + "."
			if state.NextState == nil {
				response += " You are now at final state."
			} else {
				response += " Steps:"
				for state != nil {
					response += state.GetNextMove()
					if state.NextState != nil {
						response += ", then"
					} else {
						response += "stop."
					}

					state = state.NextState
				}
			}
		}
	} else {
		response = `Hi there! This is ArbaCS6401Bot. You can ask for tasks specified bellow:
      - /task1 [position: LEFT/RIGHT] [is left clean: TRUE/FALSE] [is right clean: TRUE/FALSE]`
	}

	return response
}

func checkCleanOption(value string) bool {
	if value == "true" || value == "false" {
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
