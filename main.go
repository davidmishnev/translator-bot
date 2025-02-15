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
	botToken := "7131193852:AAGlXe5YnBRQfHsSdQ8xzR5wnY3YA1bJQ_I"
	apiKey := os.Getenv("AQVNy8rAlnAU7SxUvMb0uZqQ1rGPbPAeOhFG1lO3")
	url := "https://translate.api.cloud.yandex.net"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Authorization: ", "Api-Key "+apiKey)
	// Authorization: Api-Key AQVNy8rAlnAU7SxUvMb0uZqQ1rGPbPAeOhFG1lO3
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	opts := []bot.Option{
		bot.WithDefaultHandler(handler),
	}
	bot, err := bot.New(botToken, opts...)
	if err != nil {
		panic(err)
	}

	bot.Start(ctx)
}
