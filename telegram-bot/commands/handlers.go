package commands

import (
	"context"
	"delivery-planner-bot/services"
	"fmt"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type CommandHandler struct {
	commandStateMap map[string]Command
	b               *bot.Bot
	routeService    *services.RouteService
}

func NewCommandHandler(b *bot.Bot, routeService *services.RouteService) *CommandHandler {
	return &CommandHandler{
		commandStateMap: make(map[string]Command),
		b:               b,
		routeService:    routeService,
	}
}

func (h *CommandHandler) RegisterCommand(identifier string, cmd Command) {
	h.commandStateMap[identifier] = cmd
}

func (h *CommandHandler) HandleMessage(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message == nil {
		return
	}

	identifier := fmt.Sprintf("%d.%d", update.Message.Chat.ID, update.Message.From.ID)

	// Handle commands
	if update.Message.Text == "/start" {
		h.commandStateMap[identifier] = NewStartCommand()
	} else if update.Message.Text == "/preview" || h.isPreviewCommand(update.Message.Text) {
		if _, ok := h.commandStateMap[identifier].(*PreviewCommand); !ok {
			h.commandStateMap[identifier] = NewPreviewCommand(h.routeService)
		}
	}

	if command, ok := h.commandStateMap[identifier]; ok {
		response, err := command.Execute([]string{update.Message.Text})
		if err != nil {
			h.sendMessage(ctx, update.Message.Chat.ID, fmt.Sprintf("Error: %v", err))
			return
		}
		if response != "" {
			h.sendMessage(ctx, update.Message.Chat.ID, response)
		}
		return
	}
}

func (h *CommandHandler) isPreviewCommand(text string) bool {
	return len(text) >= len("/preview") && text[0:len("/preview")] == "/preview"
}

func (h *CommandHandler) sendMessage(ctx context.Context, chatID int64, text string) {
	_, err := h.b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: chatID,
		Text:   text,
	})
	if err != nil {
		fmt.Printf("Error sending message: %v\n", err)
	}
}
