package commands

import (
	"encoding/json"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) List(inputMessage *tgbotapi.Message) {

	outputMsgText := "Here all the products: \n\n"

	products := c.productService.List()

	for _, p := range products {
		outputMsgText += p.Title
		outputMsgText += "\n"
	}

	serializedData, _ := json.Marshal(CommandData{Offset: 999})

	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Nex Page", string(serializedData)),
		),
	)

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)
	msg.ReplyMarkup = keyboard

	c.bot.Send(msg)
}

// func init() {
// 	registeredCommands["list"] = (*Commander).List
// }
