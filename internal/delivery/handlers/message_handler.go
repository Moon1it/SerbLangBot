package handler

import (
	"fmt"

	"github.com/Moon1it/serb-lang-bot/internal/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.mongodb.org/mongo-driver/mongo"
)

type MessageResponse struct {
	Poll    tgbotapi.SendPollConfig
	Message tgbotapi.MessageConfig
}

func HandleMessage(db *mongo.Database, message *tgbotapi.Message) (MessageResponse, error) {
	switch message.Text {
	case "/start":
		message, err := service.GetStartMessage(db, message.Chat.ID, message.From.UserName)
		if err != nil {
			return MessageResponse{}, err
		}
		return MessageResponse{Message: message}, nil

	case "/menu", "Back to Menu 🏠":
		message, err := service.GetMainMenu(db, message.Chat.ID)
		if err != nil {
			return MessageResponse{}, err
		}
		return MessageResponse{Message: message}, nil

	case "Topics 📖":
		message, err := service.GetTopicMenu(db, message.Chat.ID)
		if err != nil {
			return MessageResponse{}, err
		}
		return MessageResponse{Message: message}, nil
	case "Your progress 🎯":
		message, err := service.GetUserStats(db, message.Chat.ID)
		if err != nil {
			return MessageResponse{}, err
		}
		return MessageResponse{Message: message}, nil

	default:
		pollConfig, err := service.GetExerciseByTopic(db, message.Chat.ID, message.Text)
		if err != nil {
			return MessageResponse{}, err
		}
		return MessageResponse{Poll: pollConfig}, nil
	}
}

func HandlePollAnswer(db *mongo.Database, pollAnswer *tgbotapi.PollAnswer) error {
	err := service.UpdateTopicProgress(db, pollAnswer.User.ID, pollAnswer.OptionIDs[0])
	if err != nil {
		return fmt.Errorf("failed to update topic progress: %w", err)
	}

	return nil
}
