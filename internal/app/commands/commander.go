package commands

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Pandaesp/bot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// var registeredCommands = map[string]func(c *Commander, msg *tgbotapi.Message){}

type Commander struct {
	bot            *tgbotapi.BotAPI
	productService *product.Service
}
type CommandData struct {
	Offset int `json:"offset"`
}

func NewCommander(bot *tgbotapi.BotAPI, productService *product.Service) *Commander {
	return &Commander{
		bot:            bot,
		productService: productService,
	}
}

func (c *Commander) HandleUpdate(update tgbotapi.Update) {
	defer func() {
		if panicValue := recover(); panicValue != nil {
			log.Printf("recovered from panic: %v", panicValue)
		}
	}()

	if update.CallbackQuery != nil {
		// // args := strings.Split(update.CallbackQuery.Data, "_")
		// msg := tgbotapi.NewMessage(
		// 	update.CallbackQuery.Message.Chat.ID,
		// 	fmt.Sprintf("Command: %s\n", args[0]+
		// 		fmt.Sprintf("Offset: %s\n", args[1])),
		// )

		parsrdData := CommandData{}
		json.Unmarshal([]byte(update.CallbackQuery.Data), &parsrdData)
		msg := tgbotapi.NewMessage(
			update.CallbackQuery.Message.Chat.ID,
			fmt.Sprintf("Parsed: %+v\n", parsrdData),
		)

		c.bot.Send(msg)
		return
	}

	if update.Message != nil { // If we got a message
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		// command, ok := registeredCommands[update.Message.Command()]
		// if ok {
		// 	command(c, update.Message)
		// } else {
		// 	c.Default(update.Message)
		// }

		switch update.Message.Command() {
		case "help":
			c.Help(update.Message)
		case "list":
			c.List(update.Message)
		case "get":
			c.Get(update.Message)
		default:
			c.Default(update.Message)
		}
	}
}
