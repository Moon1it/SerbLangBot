package telegram

import (
	"fmt"
	"os"

	"github.com/Moon1it/SerbLangBot/internal/handlers"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	bot *tgbotapi.BotAPI
}

func Init() (*Bot, error) {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_TOKEN"))
	if err != nil {
		return nil, fmt.Errorf("failed to initialize bot API: %w", err)
	}

	return &Bot{
		bot: bot,
	}, nil
}

func (b *Bot) Run() error {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60
	// b.bot.Debug = true
	updates := b.bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		switch {
		case update.Message != nil:
			go func(message *tgbotapi.Message) {
				response, err := handlers.HandleMessage(message)
				if err != nil {
					fmt.Println("Failed to handle message:", err)
				}

				if response.Message.Text != "" {
					b.SendMessage(response.Message)
				}

				if len(response.Poll.Options) != 0 {
					b.SendMessage(response.Poll)
				}
			}(update.Message)

		case update.PollAnswer != nil:
			go func(pollAnswer *tgbotapi.PollAnswer) {
				err := handlers.HandlePollAnswer(pollAnswer)
				if err != nil {
					fmt.Println("Failed to handle poll answer:", err)
				}
			}(update.PollAnswer)
		}
	}

	return nil
}

func (b *Bot) SendMessage(msg tgbotapi.Chattable) {
	_, err := b.bot.Send(msg)
	if err != nil {
		fmt.Println("Failed to send message:", err)
	}
}
