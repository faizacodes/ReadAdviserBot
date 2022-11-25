package bots

import (
	parser "faiza_bot/scraping"
	"fmt"
	telebot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"time"
)

var numericKeyboard = telebot.NewReplyKeyboard(
	telebot.NewKeyboardButtonRow(
		telebot.NewKeyboardButton("Вещи"),
		telebot.NewKeyboardButton("Мне не понравилось"),
		telebot.NewKeyboardButton("Беру!"),
	),
)

func BotFunc() {
	bot, err := telebot.NewBotAPI("5983787620:AAHWdWFE5A7j1l0PcOKY5vr7dX1XBaaCyMs")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := telebot.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		msg := telebot.NewMessage(update.Message.Chat.ID, update.Message.Text)

		if update.Message.IsCommand() {
			msg := telebot.NewMessage(update.Message.Chat.ID, "")
			file := telebot.NewPhoto(update.Message.From.ID, telebot.FilePath("photo.jpeg"))
			switch update.Message.Command() {
			case "start":
				msg.Text = fmt.Sprintf("🛍 Здравствуйте %s!\n"+
					"🛍 Добро пожаловать в %s\n - интернет магазин.\n 🛍 Надеюсь вы приобретете себе какую либо вещь. Хорошей покупки!",
					update.Message.From.UserName, bot.Self.UserName)
				msg.ReplyMarkup = numericKeyboard
			}
			if _, err := bot.Send(file); err != nil {
				log.Fatalln(err)
			}
			bot.Send(msg)
		}
		switch update.Message.Text {
		case "Вещи":
			for _, item := range parser.GetJsonData() {
				data := fmt.Sprintf("Name: %s \nPrice: %s \n%s", item.Name, item.Price, item.ImgUrl)
				msg.Text = data
				time.Sleep(5 * time.Second)
				bot.Send(msg)
			}
		}

	}
}
