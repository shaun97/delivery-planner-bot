package main

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(handler),
		bot.WithWebhookSecretToken(os.Getenv("EXAMPLE_TELEGRAM_WEBHOOK_SECRET_TOKEN"))
	}

	b, _ := bot.New(os.Getenv("EXAMPLE_TELEGRAM_BOT_TOKEN"), opts...)

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