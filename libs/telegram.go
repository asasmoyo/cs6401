package libs

import (
	"log"
	"net/url"
	"os"

	telegramApi "gopkg.in/telegram-bot-api.v4"
)

var bot *telegramApi.BotAPI

func init() {
	telegramToken := os.Getenv("TELEGRAM_TOKEN")
	log.Printf("Setting telegram bot using token: %s", telegramToken)

	var err error
	bot, err = telegramApi.NewBotAPI(telegramToken)
	if err != nil {
		log.Panic(err)
	}
}

// SetupBot telegram
func SetupBot(appURL string) {
	log.Printf("Detected bot account: %s", bot.Self.UserName)

	webhookURL, err := url.Parse(appURL + "/telegramHandler")
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Setting telegram webhook url to: %s", appURL)
	bot.SetWebhook(telegramApi.WebhookConfig{URL: webhookURL})
	log.Println("Done!")
}
