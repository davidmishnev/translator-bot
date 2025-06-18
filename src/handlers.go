package main

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"log"
)

func handleHelp(ctx context.Context, b *bot.Bot, update *models.Update) {
	chatID := update.Message.Chat.ID
	helpText := "Available commands:\n/start - Welcome message\n/help - List commands"

	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: chatID,
		Text:   helpText,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func handleCommand(ctx context.Context, b *bot.Bot, update *models.Update) {
	chatID := update.Message.Chat.ID
	command := update.Message.Text
	switch command {
	case "/help":
		handleHelp(ctx, b, update)
	case "/start":
		_, err := b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chatID,
			Text:   "pls send a picture/text",
		})
		if err != nil {
			log.Println("Error sending /start response:", err)
		}
	default:
		_, err := b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chatID,
			Text:   "Unknown command: " + command + ".\nTry out /help",
		})
		if err != nil {
			log.Println("Error sending unknown command response:", err)
		}
	}
}

func messageHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message.Photo != nil && len(update.Message.Photo) > 0 {
		handlePhoto(ctx, b, update)
		return
	}

	if update.Message.Text != "" {
		handleText(ctx, b, update)
		return
	}

	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Please send a photo or text message.",
	})
	if err != nil {
		log.Println("Error sending message:", err)
	}
}

func handleText(ctx context.Context, b *bot.Bot, update *models.Update) {
	chatID := update.Message.Chat.ID
	text := update.Message.Text

	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: chatID,
		Text:   text,
	})
	if err != nil {
		log.Println("Error sending text response:", err)
	}
}

// todo: finish the photo handler
func handlePhoto(ctx context.Context, b *bot.Bot, update *models.Update) {
	chatID := update.Message.Chat.ID
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: chatID,
		Text:   "qwe",
	})
	if err != nil {
		log.Println("Error sending photo:", err)
	}
}
