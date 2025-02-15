package main

import (
	"log"
	"net/http"
	"os"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
)

func main() {
	botToken := "7131193852:AAGlXe5YnBRQfHsSdQ8xzR5wnY3YA1bJQ_I"
	apiKey := os.Getenv("AQVNy8rAlnAU7SxUvMb0uZqQ1rGPbPAeOhFG1lO3")
	url := "https://translate.yandex.com/"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Authorization: ", "Api-Key "+apiKey)
	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())
	if err != nil {
		log.Fatal(err)
	}
	_ = apiKey
	// AQVNy8rAlnAU7SxUvMb0uZqQ1rGPbPAeOhFG1lO3
	updates, _ := bot.UpdatesViaLongPolling(nil)
	defer bot.StopLongPolling()
	bh, _ := th.NewBotHandler(bot, updates)
	defer bh.Stop()
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		if update.Message != nil && update.Message.Text == "PSG" || update.Message.Text == "psg" || update.Message.Text == "псж" {
			chatID := update.Message.Chat.ID
			inlineKeyboard := &telego.InlineKeyboardMarkup{
				InlineKeyboard: [][]telego.InlineKeyboardButton{
					{
						{Text: "Псж красное", CallbackData: "btn_01", URL: "https://fitp.itmo.ru/files/lgd.pdf"},
						{Text: "Псж синее", CallbackData: "btn_02", URL: "https://fitp.itmo.ru/files/lgd.pdf"},
					},
				},
			}
			_, _ = bot.SendMessage(&telego.SendMessageParams{
				ChatID:      telego.ChatID{ID: chatID},
				Text:        "Выбери свое псж",
				ReplyMarkup: inlineKeyboard,
			})
		}
	})
	bh.Start()
	defer bh.Stop()

}
