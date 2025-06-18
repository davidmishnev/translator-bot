package main

import (
	"context"
	"github.com/go-telegram/bot"
	"log"
	"net/http"
	"os"
	"os/signal"
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

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	BOT, err := bot.New(botToken)
	BOT.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, handleCommand)
	BOT.RegisterHandler(bot.HandlerTypeMessageText, "/help", bot.MatchTypeExact, handleHelp)
	//BOT.RegisterHandler(bot.HandlerTypeCallbackQueryData, "button0", 0, )
	BOT.RegisterHandler(bot.HandlerTypeMessageText, "", bot.MatchTypeContains, messageHandler)

	BOT.Start(ctx)
}
