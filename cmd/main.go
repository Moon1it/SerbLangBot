package main

import (
	"log"

	"github.com/Moon1it/SerbLangBot/internal/clients/telegram"
	"github.com/Moon1it/SerbLangBot/internal/handler"
	"github.com/Moon1it/SerbLangBot/internal/repository"
	"github.com/Moon1it/SerbLangBot/internal/service"
	"github.com/Moon1it/SerbLangBot/pkg/database"
)

func main() {
	// Load environment variables from the .env file
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Failed to load .env file:", err)
	// }

	// Connect to the MongoDB database
	mongoDB, err := database.ConnectToMongoDB()
	if err != nil {
		log.Fatal("Error starting MongoDB:", err)

	}
	// Defer the closing of the MongoDB connection
	defer func() {
		err := mongoDB.CloseMongoDB()
		if err != nil {
			log.Fatal("Error closing MongoDB:", err)
		}
	}()

	repo := repository.InitRepository(mongoDB.Client)
	service := service.InitService(repo)
	handler := handler.InitHandler(service)

	bot, err := telegram.Init(handler)
	if err != nil {
		log.Fatalf("Failed to initialize Telegram bot: %v", err)
	}

	if err := bot.Run(); err != nil {
		log.Fatalf("Failed to run the bot: %v", err)
	}
}
