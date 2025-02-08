package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(handler),
		bot.WithWebhookSecretToken(os.Getenv("TELEGRAM_TOKEN")),
	}

	b, _ := bot.New(os.Getenv("TELEGRAM_TOKEN"), opts...)

	// call methods.SetWebhook if needed

	go b.StartWebhook(ctx)

	http.ListenAndServe(":2000", b.WebhookHandler())

	// call methods.DeleteWebhook if needed
}

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   update.Message.Text,
	})
}
