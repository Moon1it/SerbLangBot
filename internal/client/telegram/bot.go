package telegram

import (
	"fmt"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	BotAPI      *tgbotapi.BotAPI
	UpdatesChan tgbotapi.UpdatesChannel
}

func Init() (*Bot, error) {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_TOKEN"))
	if err != nil {
		return nil, fmt.Errorf("failed to initialize bot API: %w", err)
	}
	bot.Debug = true

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updates := bot.GetUpdatesChan(updateConfig)

	return &Bot{
		BotAPI:      bot,
		UpdatesChan: updates,
	}, nil
}

func (b *Bot) SendMessage(msg tgbotapi.Chattable) {
	_, err := b.BotAPI.Send(msg)
	if err != nil {
		fmt.Println("Failed to send message:", err)
	}
}
