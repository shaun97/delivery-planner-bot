package main

import (
	"context"
	"fmt"
	"myproject/bot"
	"myproject/config"
	"os"
)

func main() {
	err := config.LoadEnv()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	b, err := bot.Initialize(os.Getenv("TELEGRAM_TOKEN"))
	if err != nil {
		fmt.Println(err)
	}

	b.Start(context.TODO())
}
