package handlers

import (
	"fmt"

	"github.com/Moon1it/SerbLangBot/internal/services"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type MessageResponse struct {
	Poll    tgbotapi.SendPollConfig
	Message tgbotapi.MessageConfig
}

func HandleMessage(message *tgbotapi.Message) (MessageResponse, error) {
	switch message.Text {
	case "/start":
		message, err := services.GetStartMessage(message.Chat.ID, message.From.UserName)
		if err != nil {
			return MessageResponse{}, err
		}
		return MessageResponse{Message: message}, nil

	case "/menu", "Back to Menu 🏠":
		message, err := services.GetMainMenu(message.Chat.ID)
		if err != nil {
			return MessageResponse{}, err
		}
		return MessageResponse{Message: message}, nil

	case "Topics 📖":
		message, err := services.GetTopicMenu(message.Chat.ID)
		if err != nil {
			return MessageResponse{}, err
		}
		return MessageResponse{Message: message}, nil
	case "Your progress 🎯":
		message, err := services.GetUserProgress(message.Chat.ID)
		if err != nil {
			return MessageResponse{}, err
		}
		return MessageResponse{Message: message}, nil

	default:
		pollConfig, err := services.GetExerciseByTopic(message.Chat.ID, message.Text)
		if err != nil {
			return MessageResponse{}, err
		}
		return MessageResponse{Poll: pollConfig}, nil
	}
}

func HandlePollAnswer(pollAnswer *tgbotapi.PollAnswer) error {
	err := services.UpdateTopicProgress(pollAnswer.User.ID, pollAnswer.OptionIDs[0])
	if err != nil {
		return fmt.Errorf("failed to update topic progress: %w", err)
	}

	return nil
}
