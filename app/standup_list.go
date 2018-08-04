package app

import (
	"github.com/andrysds/clarity"
	"github.com/andrysds/panera/entity"
	"gopkg.in/telegram-bot-api.v4"
)

func (p *Panera) HandleStandupList(update *tgbotapi.Update) {
	standups, err := entity.GetStandupList()
	clarity.PrintIfError(err, "error on get standup list")

	messageText := "Stand up lead periode ini:\n"
	for _, s := range standups {
		if s.State == "1" {
			messageText += "`[x]` "
		} else {
			messageText += "`[ ]` "
		}
		messageText += s.Name + "\n"
	}

	message := p.NewMessage(update.Message.Chat.ID, messageText)
	p.SendMessage(message)
}