package commands

import (
	"example.com/bot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Commander struct {
	bot            *tgbotapi.BotAPI
	productService *product.Service
}

func NewCommander(productService *product.Service,
	bot *tgbotapi.BotAPI) *Commander {
	return &Commander{
		bot:            bot,
		productService: productService,
	}
}
