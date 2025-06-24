package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-telegram/bot"
	"io"
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

type ResponseData struct {
	Text string `json:"text"`
}

var jsonValue jsonData
var apiKey = os.Getenv("AQVNyxkKoJu7PK42eAXJKMi7AfwBuKn3HluCWw08")        //api
var url = "https://translate.api.cloud.yandex.net/translate/v2/translate" //url

func saveToJson() {
	data, err := json.MarshalIndent(jsonValue, "", "  ")
	if err != nil {
		fmt.Println("Ошибка сериализации:", err)
		return
	}
	err = os.WriteFile("body.json", data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func makeRequest() (string, error) {
	jsonBytes, err := os.ReadFile("body.json")
	if err != nil {
		log.Fatal(err)
	}
	buffer := bytes.NewBuffer(jsonBytes)
	req, err := http.NewRequest("POST", url, buffer)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Authorization", "Api-Key "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var result ResponseData
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatal(err)
	}

	var translatedText string
	if result.Text != "" {
		translatedText = result.Text
	} else {
		translatedText = "перевод не найден"
	}
	return translatedText, nil
}

func main() {

	botToken := "7131193852:AAGlXe5YnBRQfHsSdQ8xzR5wnY3YA1bJQ_I" //bot

	folderIDvalue := "b1gtccjhdqmvpftp1ctc"

	jsonValue.FolderId = folderIDvalue

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
