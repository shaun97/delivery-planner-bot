package main

import (
	"context"
	"fmt"
	"os"

	"delivery-planner-bot/commands"
	"delivery-planner-bot/services"

	"github.com/go-telegram/bot"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	token := os.Getenv("TELEGRAM_TOKEN")
	baseURL := os.Getenv("BACKEND_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8080"
	}

	b, err := bot.New(token)
	if err != nil {
		fmt.Printf("Error creating bot: %v\n", err)
		os.Exit(1)
	}

	routeService := services.NewRouteService(baseURL)

	commandHandler := commands.NewCommandHandler(b, routeService)

	// Register handlers for all messages
	b.RegisterHandler(bot.HandlerTypeMessageText, "", bot.MatchTypePrefix, commandHandler.HandleMessage)

	fmt.Println("Bot started...")
	b.Start(context.Background())
}
