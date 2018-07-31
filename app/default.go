package app

import "gopkg.in/telegram-bot-api.v4"

func (p *Panera) HandleDefault(update *tgbotapi.Update) {
	messageText := "Hello World! Panera here."
	message := tgbotapi.NewMessage(update.Message.Chat.ID, messageText)
	p.SendMessage(message)
}