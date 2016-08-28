package libs

import (
	telegramApi "gopkg.in/telegram-bot-api.v4"
	"log"
	"net/url"
)

// SetupBot telegram
func SetupBot(token string, appURL string) {
	log.Printf("Setting up telegram bot using token: %s", token)
	bot, error := telegramApi.NewBotAPI(token)
	if error != nil {
		log.Panic(error)
	}

	log.Printf("Detected bot account: %s", bot.Self.UserName)

	webhookURL, error := url.Parse(appURL + "/telegramHandler")
	if error != nil {
		log.Panic(error)
	}

	log.Println("Setting up webhook url...")
	bot.SetWebhook(telegramApi.WebhookConfig{URL: webhookURL})
	log.Println("Done!")
}
