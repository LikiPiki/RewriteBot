package main

import (
	"gopkg.in/telegram-bot-api.v4"
)

const (
	TOKEN   = "119720842:AAEJniFGSN4d3_6PeC6bdJxY_hWcy444ZM8"
	TIMEOUT = 60
)

var (
	bot *tgbotapi.BotAPI

	sendFlag = true
)

func main() {
	var err error
	bot, err = tgbotapi.NewBotAPI(TOKEN)

	if err != nil {
		panic(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = TIMEOUT
	updates, err := bot.GetUpdatesChan(u)

	if err != nil {
		panic(err)
	}
	for update := range updates {
		msg := update.Message
		if msg == nil {
			continue
		}

		if msg.IsCommand() {
			command := msg.Command()
			switch command {
			case "kek":
				go generatePhrase(msg, "kek")
			case "or":
				go generatePhrase(msg, "or")
			case "info":
				go generatePhrase(msg, "info")
			case "telki":
				go sendNudes(msg)
			case "turn":
				go turnFlag(msg)
			}
		} else {
			go generatePhrase(msg, "sendRofl")
		}

	}
}
