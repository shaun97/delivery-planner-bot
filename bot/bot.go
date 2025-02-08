package bot

import (
	"github.com/go-telegram/bot"
	"myproject/bot/handlers"
)

func Initialize(token string) (*bot.Bot, error) {
	return bot.New(token, bot.WithDefaultHandler(handlers.DefaultHandler))
}
