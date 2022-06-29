package msgs

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func CreateNewTGMessage(text string) tgbotapi.Chattable {
	message := tgbotapi.MessageConfig{
		BaseChat: tgbotapi.BaseChat{
			ChatID: 872383555,
		},
		Text: text,
	}

	return message
}

func BotSendMsg(bot *tgbotapi.BotAPI, msg tgbotapi.Chattable) {
	_, err := bot.Send(msg)
	if err != nil {
		log.Println(err)
	}
}
