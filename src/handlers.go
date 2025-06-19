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
			Text:   "Привет, это бот переводчик, пожалуйста, отправьте текст который вы хотите перевести!",
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
	jsonValue.Texts = update.Message.Text
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "На какой язык вы хотите перевести этот текст? ",
		ReplyMarkup: &models.InlineKeyboardMarkup{
			InlineKeyboard: [][]models.InlineKeyboardButton{
				{
					{
						Text:         "Русский",
						CallbackData: "button_ru",
					},
					{
						Text:         "Английский",
						CallbackData: "button_en",
					},
					{
						Text:         "Французский",
						CallbackData: "button_fr",
					},
					{
						Text:         "Японский",
						CallbackData: "button_ja",
					},
				},
			},
		},
	})
	if err != nil {
		log.Println("Ошибка отправки сообщения:", err)
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

func handleCallbackData(ctx context.Context, b *bot.Bot, update *models.Update) {
	callbackData := update.CallbackQuery.Data
	switch callbackData {
	case "button_ru":
		jsonValue.TargetLanguageCode = "ru"
		break
	case "button_en":
		jsonValue.TargetLanguageCode = "en"
		break
	case "button_fr":
		jsonValue.TargetLanguageCode = "fr"
		break
	case "button_ja":
		jsonValue.TargetLanguageCode = "ja"
		break
	default:
		break
	}
}
