package main

import (
	"context"
	"github.com/go-telegram/bot"
	"log"
	"net/http"
	"os"
)

func main() {

	botToken := "7131193852:AAGlXe5YnBRQfHsSdQ8xzR5wnY3YA1bJQ_I"    //bot
	apiKey := os.Getenv("AQVNy8rAlnAU7SxUvMb0uZqQ1rGPbPAeOhFG1lO3") //api
	// AQVNy8rAlnAU7SxUvMb0uZqQ1rGPbPAeOhFG1lO3
	url := "https://translate.api.cloud.yandex.net" //url

	req, err := http.NewRequest("GET", url, nil) //http request
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Authorization: ", "Api-Key "+apiKey) //Request header

	BOT, err := bot.New(botToken)

	BOT.RegisterHandler(bot.HandlerTypeMessageText, "/", 0, handleTextMessage)
	BOT.RegisterHandler(bot.HandlerTypeMessageText, "/help", 0, handleHelp)
	BOT.RegisterHandler(bot.HandlerTypeMessageText, "", 0, handler)
	BOT.Start(context.Background())
}
