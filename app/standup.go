package app

import (
	"fmt"

	"github.com/andrysds/clarity"
	"github.com/andrysds/panera/entity/standup"
	"gopkg.in/telegram-bot-api.v4"
)

func (p *Panera) HandleStandup(update *tgbotapi.Update) {
	standup, err := standup.Current()
	clarity.PrintIfError(err, "error on get standup")

	messageTemplate := "Yuk stand up! Yang dapat giliran untuk memimpin stand up hari ini adalah _%s_ (@%s)"
	messageText := fmt.Sprintf(messageTemplate, standup.Name, standup.Username)

	message := p.NewMessage(update.Message.Chat.ID, messageText)
	p.SendMessage(message)
}

func (p *Panera) HandleStandupSkip(update *tgbotapi.Update) {
	standup, current, err := standup.Next()
	clarity.PrintIfError(err, "error on skipping standup")

	if err == nil {
		messageTemplate := "Karena %s tidak bisa, penggantinya _%s_ (@%s)"
		messageText := fmt.Sprintf(messageTemplate, current.Name, standup.Name, standup.Username)
		message := p.NewMessage(update.Message.Chat.ID, messageText)
		p.SendMessage(message)
	} else if err.Error() == "not found" {
		messageText := "Waduh ga ada gantinya lagi nih!"
		message := p.NewMessage(update.Message.Chat.ID, messageText)
		p.SendMessage(message)
	}
}
