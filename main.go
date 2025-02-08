package main

import (
	"context"
	"os"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func main() {
	b, _ := bot.New((os.Getenv("TELEGRAM_TOKEN")), bot.WithDefaultHandler(handler))

	b.Start(context.TODO())
}

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	print("hello world")
}
