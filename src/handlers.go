package main

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"log"
)

func handleTextMessage(ctx context.Context, b *bot.Bot, update *models.Update) {
	chatID := update.Message.Chat.ID
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: chatID,
		Text:   "You sent: " + update.Message.Text,
	})
	if err != nil {
		log.Fatal(err)
	}
}

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
	if update.Message == nil {
		return
	}

	chatID := update.Message.Chat.ID
	command := update.Message.Text

	switch command {
	case "/help":
		handleHelp(ctx, b, update) //isn't it redundant? could be

	case "/start":
		_, err := b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chatID,
			Text:   "Please, enter output message language",
		})
		if err != nil {
			log.Println("Error sending /start response:", err)
		}

	default:
		_, err := b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chatID,
			Text:   "Unknown command: " + command + ". Try /help",
		})
		if err != nil {
			log.Println("Error sending unknown command response:", err)
		}
	}
}

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message == nil {
		return
	}

	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   update.Message.Text,
	})
	if err != nil {
		log.Fatal(err)
	}
}
