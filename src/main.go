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

	botToken := "TOKEN"    //bot
	apiKey := os.Getenv("TOKEN") //api
	url := "https://translate.api.cloud.yandex.net" //url

	req, err := http.NewRequest("GET", url, nil) //http request
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Authorization: ", "Api-Key "+apiKey) //Request header

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	opts := []bot.Option{
		bot.WithDefaultHandler(handleCommand),
	}
	BOT, err := bot.New(botToken, opts...)
	BOT.Start(ctx)
}
