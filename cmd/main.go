package main

import (
	"log"

	"github.com/Moon1it/serb-lang-bot/internal/app"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env file", err)
	}

	app, err := app.Init()
	if err != nil {
		log.Fatal("Error start app", err)
	}
	app.Run()
}
