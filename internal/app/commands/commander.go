package commands

import (
	"log"

	"github.com/Pandaesp/bot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var registeredCommands = map[string]func(c *Commander, msg *tgbotapi.Message){}

type Commander struct {
	bot            *tgbotapi.BotAPI
	productService *product.Service
}

func NewCommander(bot *tgbotapi.BotAPI, productService *product.Service) *Commander {
	return &Commander{
		bot:            bot,
		productService: productService,
	}
}

func (c *Commander) HandleUpdate(update tgbotapi.Update) {
	if update.Message != nil { // If we got a message
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		command, ok := registeredCommands[update.Message.Command()]
		if ok {
			command(c, update.Message)
		} else {

		}
		// switch update.Message.Command() {
		// case "help":
		// 	c.Help(update.Message)
		// case "list":
		// 	c.List(update.Message)
		// default:
		// 	c.Default(update.Message)
		// }
	}
}
