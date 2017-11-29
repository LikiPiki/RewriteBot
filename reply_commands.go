package main

import (
	"gopkg.in/telegram-bot-api.v4"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
)

var (
	api = "http://api.oboobs.ru/noise/1/"
)

func generatePhrase(msg *tgbotapi.Message, mode string) {
	var reply tgbotapi.MessageConfig
	switch mode {
	case "kek":
		reply = tgbotapi.NewMessage(msg.Chat.ID, keks[rand.Intn(len(keks))])
	case "or":
		reply = tgbotapi.NewMessage(msg.Chat.ID, "Ору "+ornul[rand.Intn(len(ornul))])
	case "info":
		text := "*Админы*:\n"
		for _, i := range admins {
			text += fmt.Sprintf("%s\n", i)
		}
		reply = tgbotapi.NewMessage(msg.Chat.ID, text)
	case "error":
		reply = tgbotapi.NewMessage(msg.Chat.ID, errors[rand.Intn(len(errors))])
	case "changeFlag":
		var text string
		if sendFlag {
			text = "Ура, *ПОГНАЛЛЛЛИИИИ*, что там по телочкам!"
		} else {
			reply = tgbotapi.NewMessage(msg.Chat.ID, telki_ne_daut[rand.Intn(len(telki_ne_daut))])
			bot.Send(reply)
			text = "*Телки* отключены!("
		}
		reply = tgbotapi.NewMessage(msg.Chat.ID, text)
	case "sendRofl":
		if rand.Intn(2) == 1 {
			st := tgbotapi.NewStickerShare(msg.Chat.ID, stikers[rand.Intn(len(stikers))])
			bot.Send(st)
		} else {
			reply = tgbotapi.NewMessage(msg.Chat.ID, replys[rand.Intn(len(replys))])
		}
	default:
		reply = tgbotapi.NewMessage(msg.Chat.ID, mode)
	}
	reply.ParseMode = "markdown"
	bot.Send(reply)
}

func sendNudes(msg *tgbotapi.Message) {
	if sendFlag {
		resp, err := http.Get(api)
		defer resp.Body.Close()
		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			generatePhrase(msg, "error")
		}
		var q Querys
		err = json.Unmarshal(bytes, &q)
		if err != nil {
			println(err.Error())
			generatePhrase(msg, "error")
		} else {
			photo := tgbotapi.NewPhotoShare(msg.Chat.ID, fmt.Sprintf("http://media.oboobs.ru/noise/%d.jpg", q[0].Id))
			bot.Send(photo)
		}
	} else {
		generatePhrase(msg, "*Запрещено* админами!")
	}
}

func turnFlag(msg *tgbotapi.Message) {
	fl := false
	for _, e := range admins {
		if e == msg.From.UserName {
			fl = true
			break
		}
	}
	if fl {
		sendFlag = !sendFlag
		generatePhrase(msg, "changeFlag")
	} else {
		generatePhrase(msg, "Ты не админ, хочешь в илиту? Спроси меня как @likipiki")
	}
}
