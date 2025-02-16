package main

import (
	"context"
	"delivery-planner-bot/commands"
	"fmt"
	"os"

	"github.com/go-telegram/bot"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	commandHandler := commands.NewCommandHandler()

	b, err := bot.New((os.Getenv("TELEGRAM_TOKEN")), bot.WithDefaultHandler(commandHandler.HandleMessage))
	if err != nil {
		fmt.Println(err)
	}

	b.Start(context.TODO())
}
