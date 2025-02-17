package main

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"log"
)

func handleTextMessage(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message != nil {
		return
	}
	chatID := update.Message.Chat.ID
	language := update.Message.Text
	if isValidLanguage(language) {
		//json -> language
	} else {
		_, err := b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chatID,
			Text:   "Please enter a valid language",
		})
		if err != nil {
			log.Println(err)
		}
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
		handleHelp(ctx, b, update)
	case "/start":
		_, err := b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chatID,
			Text:   "Please, enter output message language",
		})
		if err != nil {
			log.Println("Error sending /start response:", err)
		}
		b.RegisterHandler(bot.HandlerTypeMessageText, "", 0, handleTextMessage)
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
