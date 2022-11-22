package main

import (
	"log"
	"fmt"
	"time"


	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)
var numericKeyboard1 = tgbotapi.NewReplyKeyboard(
    tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Привет"),
		tgbotapi.NewKeyboardButton("Как ты?"),
		tgbotapi.NewKeyboardButton("Чем занимаешься?"),
		tgbotapi.NewKeyboardButton("Сегодняшняя дата?"),
),
)

var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
    tgbotapi.NewInlineKeyboardRow(
        tgbotapi.NewInlineKeyboardButtonData("Хорошо", "Рада знать! У меня тоже все хорошо"),
        tgbotapi.NewInlineKeyboardButtonData("Плохо", "Жаль, надеюсь тебе станет лучше"),
        tgbotapi.NewInlineKeyboardButtonData("Нормально", "Хорошо что это так "),
    ),
    tgbotapi.NewInlineKeyboardRow(
        tgbotapi.NewInlineKeyboardButtonData("Так себе", "Все наладится"),
        tgbotapi.NewInlineKeyboardButtonData("Замечательно!", "Ого!"),
        tgbotapi.NewInlineKeyboardButtonData("...", "Что такое? Не хочешь говорить?"),
    ),
)


func main() {
	bot, err := tgbotapi.NewBotAPI("5632596572:AAH1X7JWrcihGi6JXgLblNn_uScDVEXlkFw")
	if err != nil {
		log.Panic(err)
	}

    bot.Debug = true

    log.Printf("Authorized on account %s", bot.Self.UserName)

    u := tgbotapi.NewUpdate(0)
    u.Timeout = 60

    updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { 
			continue
		}
	if update.Message.IsCommand() {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		file := tgbotapi.NewPhoto(update.Message.From.ID, tgbotapi.FilePath("Durak.webp"))
		switch update.Message.Command() {
		case "start":
		 msg.Text = fmt.Sprintf("Привет %s! "+
		  "Добро пожаловать в %s\nВы можете поговорите с ботом. Удачи!",
		  update.Message.From.UserName, bot.Self.UserName)
		 msg.ReplyMarkup = numericKeyboard1
		}
		if _, err := bot.Send(file); err != nil {
		 log.Fatalln(err)
		}
		bot.Send(msg)
	   }

	for update := range updates {
		if update.Message == nil { 
			continue
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

	    switch update.Message.Text {
	    case "close":
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	   }
   
	   if _, err := bot.Send(msg); err != nil {
		   log.Panic(err)
	   }

    for update := range updates {
        if update.Message != nil {
            msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
            switch update.Message.Text {
            case "Привет":
                msg.ReplyMarkup = numericKeyboard
				msg.Text = "Привет! Как ты себя чувствуешь?"
			case "Как ты?":
				msg.Text = "У меня все хорошо! Спасибо, что поинтересовалcя(-ась)"
			case "Чем занимаешься?":
				msg.Text = "Общаюсь с тобой, дел у меня больше то и нет.."
			case "Сегодняшняя дата?":
				t := time.Now()
				formatted := fmt.Sprintf("%d-%02d-%02d  %02d:%02d",
				t.Year(), t.Month(), t.Day(),
				t.Hour(), t.Minute())
				msg.Text = formatted
            }


            if _, err = bot.Send(msg); err != nil {
                panic(err)
            }
        } else if update.CallbackQuery != nil {
            
            callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
            if _, err := bot.Request(callback); err != nil {
                panic(err)
            }


            msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)
            if _, err := bot.Send(msg); err != nil {
                panic(err)
            }
			
        }
    }
	}
}
}