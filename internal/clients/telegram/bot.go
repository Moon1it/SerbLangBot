package telegram

import (
	"fmt"
	"os"

	"github.com/Moon1it/SerbLangBot/internal/handler"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	bot     *tgbotapi.BotAPI
	handler handler.Handler
}

func Init(handler *handler.Handler) (*Bot, error) {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_TOKEN"))
	if err != nil {
		return nil, fmt.Errorf("failed to initialize bot API: %w", err)
	}

	return &Bot{
		bot:     bot,
		handler: *handler,
	}, nil
}

func (b *Bot) Run() error {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60
	b.bot.Debug = true
	updates := b.bot.GetUpdatesChan(updateConfig)

	for update := range updates {

		if update.Message != nil {
			go func(message *tgbotapi.Message) {
				response, err := b.handler.HandleMessage(message)
				if err != nil {
					fmt.Println("Failed to handle message:", err)
				}

				if response.Message != nil && response.Message.Text != "" {
					_, err := b.bot.Send(response.Message)
					if err != nil {
						fmt.Println("Failed to send message:", err)
					}
					return
				}

				if response.Poll != nil && len(response.Poll.Options) != 0 {
					_, err := b.bot.Send(response.Poll)
					if err != nil {
						fmt.Println("Failed to send message:", err)
					}
				}
			}(update.Message)
		}

		if update.PollAnswer != nil {
			go func(pollAnswer *tgbotapi.PollAnswer) {
				err := b.handler.HandlePollAnswer(pollAnswer)
				if err != nil {
					fmt.Println("Failed to handle poll answer:", err)
				}
			}(update.PollAnswer)
		}
	}

	return nil
}
