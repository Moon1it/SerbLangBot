package main

import (
	"log"

	"github.com/Moon1it/SerbLangBot/internal/clients/telegram"
	"github.com/Moon1it/SerbLangBot/pkg/database"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env file:", err)
	}

	err = database.StartMongoDB()
	if err != nil {
		log.Fatal("Error starting MongoDB:", err)

	}
	defer func() {
		err := database.CloseMongoDB()
		if err != nil {
			log.Fatal("Error closing MongoDB:", err)
		}
	}()

	bot, err := telegram.Init()
	if err != nil {
		log.Fatalf("Failed to initialize Telegram bot: %v", err)
	}

	if err := bot.Run(); err != nil {
		log.Fatalf("Failed to run the bot: %v", err)
	}
}
