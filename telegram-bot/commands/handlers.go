package commands

import (
	"context"
	"fmt"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type CommandHandler struct {
	commandStateMap map[string]Command
}

func NewCommandHandler() *CommandHandler {
	return &CommandHandler{
		commandStateMap: make(map[string]Command),
	}
}

func (h *CommandHandler) HandleMessage(ctx context.Context, b *bot.Bot, update *models.Update) {
	identifier := fmt.Sprintf("%d.%d", update.Message.Chat.ID, update.Message.From.ID)
	switch update.Message.Text {
	case "/start":
		// Handle start command
		h.commandStateMap[identifier] = NewStartCommand()
	case "/help":
		// Handle help command
		print("Help command received")
	default:
		print("hello world")
	}

	if command, ok := h.commandStateMap[identifier]; ok {
		response, err := command.Execute([]string{update.Message.Text})
		if err != nil {
			print(err)
		}
		print(response)
		return
	}
	print(update.Message.Text)
}
