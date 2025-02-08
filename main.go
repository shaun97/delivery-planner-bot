package main

import (
	"context"
	"fmt"
	"os"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	b, err := bot.New((os.Getenv("TELEGRAM_TOKEN")), bot.WithDefaultHandler(handler))
	if err != nil {
		fmt.Println(err)
	}

	b.Start(context.TODO())
}

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	print("hello world")
}
