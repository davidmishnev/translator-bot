package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-telegram/bot"
	"log"
	"net/http"
	"os"
	"os/signal"
)

type jsonData struct {
	FolderId           string `json:"folderId"`
	Texts              string `json:"texts"`
	TargetLanguageCode string `json:"targetLanguageCode"`
}

var jsonValue jsonData

func main() {

	botToken := "7131193852:AAGlXe5YnBRQfHsSdQ8xzR5wnY3YA1bJQ_I"    //bot
	apiKey := os.Getenv("AQVNyxkKoJu7PK42eAXJKMi7AfwBuKn3HluCWw08") //api
	url := "https://translate.api.cloud.yandex.net"                 //url
	folderIDvalue := "b1gtccjhdqmvpftp1ctc"

	jsonValue.FolderId = folderIDvalue
	data, err := json.MarshalIndent(jsonValue, "", "  ")
	if err != nil {
		fmt.Println("Ошибка сериализации:", err)
		return
	}
	err = os.WriteFile("../body.json", data, 0644)
	if err != nil {
		fmt.Println("Ошибка записи:", err)
		return
	}
	req, err := http.NewRequest("GET", url, nil) //http request
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Authorization: ", "Api-Key "+apiKey) //Request header

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	BOT, err := bot.New(botToken)
	if err != nil {
		log.Fatal(err)
	}
	BOT.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, handleCommand)
	BOT.RegisterHandler(bot.HandlerTypeMessageText, "/help", bot.MatchTypeExact, handleHelp)
	BOT.RegisterHandler(bot.HandlerTypeMessageText, "", bot.MatchTypeContains, messageHandler)
	BOT.RegisterHandler(bot.HandlerTypeCallbackQueryData, "button_", bot.MatchTypePrefix, handleCallbackData)
	BOT.Start(ctx)
}
