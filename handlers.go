package main

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/mymmrac/telego"
	"github.com/mymmrac/telego/telegoutil"
)

func handleTextMessage(bot *telego.Bot, update telego.Update) {
	chatID := update.Message.Chat.ChatID()
	_, _ = bot.SendMessage(telegoutil.Message(chatID, "You sent: "+update.Message.Text))
}

func handleHelp(bot *telego.Bot, update telego.Update) {
	chatID := update.Message.Chat.ChatID()
	helpText := "Available commands:\n/start - Welcome message\n/help - List commands"
	_, _ = bot.SendMessage(telegoutil.Message(chatID, helpText))
}

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   update.Message.Text,
	})
}
