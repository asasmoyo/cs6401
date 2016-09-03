package telegram

import (
	"log"
	"net/url"

	telegramApi "gopkg.in/telegram-bot-api.v4"
)

// Bot struct
type Bot struct {
	instance *telegramApi.BotAPI
}

// NewBot create a new bot
func NewBot(telegramToken, appURL string) Bot {
	log.Printf("Creating bot instance using token: %s", telegramToken)
	instance, err := telegramApi.NewBotAPI(telegramToken)
	if err != nil {
		log.Panic(err)
	}

	rawWebhookURL := appURL + "/telegramHandler"
	log.Printf("Setting up webhook url to: %s", rawWebhookURL)
	webhookURL, err := url.Parse(rawWebhookURL)
	if err != nil {
		log.Panic(err)
	}
	instance.SetWebhook(telegramApi.WebhookConfig{URL: webhookURL})

	bot := Bot{instance}
	return bot
}

// SendMessage send message
func (bot Bot) SendMessage(chatID int64, textMessage string) {
	chatMessage := telegramApi.NewMessage(chatID, textMessage)
	bot.instance.Send(chatMessage)
}
