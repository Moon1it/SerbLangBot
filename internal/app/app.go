package app

import (
	"fmt"

	"github.com/Moon1it/SerbLangBot/internal/client/telegram"
	"github.com/Moon1it/SerbLangBot/internal/config"
	handler "github.com/Moon1it/SerbLangBot/internal/delivery/handlers"
	"github.com/Moon1it/SerbLangBot/pkg/clients/mongodb"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	cfg *config.Config
	db  *mongo.Database
	bot *telegram.Bot
}

func Init() (*App, error) {
	cfg, err := config.Init()
	if err != nil {
		err = fmt.Errorf("failed to create config: %w", err)
		return nil, err
	}

	bot, err := telegram.Init()
	if err != nil {
		err = fmt.Errorf("failed to create bot: %w", err)
		return nil, err
	}

	db, err := mongodb.ConnectToMongoDB(cfg)
	if err != nil {
		err = fmt.Errorf("failed to create database client: %w", err)
		return nil, err
	}

	return &App{
		cfg: cfg,
		db:  db,
		bot: bot,
	}, nil
}

func (app *App) Run() {
	for update := range app.bot.UpdatesChan {
		switch {
		case update.Message != nil:
			go func(message *tgbotapi.Message) {
				response, err := handler.HandleMessage(app.db, message)
				if err != nil {
					fmt.Println("Failed to handle message:", err)
				}

				if response.Message.Text != "" {
					app.bot.SendMessage(response.Message)
				}

				if len(response.Poll.Options) != 0 {
					app.bot.SendMessage(response.Poll)
				}
			}(update.Message)

		case update.PollAnswer != nil:
			err := handler.HandlePollAnswer(app.db, update.PollAnswer)
			if err != nil {
				fmt.Println("Failed to handle poll answer:", err)
				continue
			}
		}
	}
}
