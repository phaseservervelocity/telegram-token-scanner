package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN environment variable not set")
	}

	b, err := bot.New(token)
	if err != nil {
		log.Fatal(err)
	}

	scanner := NewScanner()
	b.RegisterHandler(bot.HandlerTypeMessageText, "", bot.MatchTypePrefix, scanner.HandleMessage)

	log.Println("Starting Telegram Token Scanner bot...")
	b.Start(ctx)
}