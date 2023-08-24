package handler

import (
	"errors"
	"fmt"

	"github.com/Moon1it/SerbLangBot/internal/models"
	"github.com/Moon1it/SerbLangBot/internal/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Handler struct {
	services *service.Service
}

func InitHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) HandleMessage(message *tgbotapi.Message) (*models.MessageResponse, error) {
	switch message.Text {
	case "/start":
		_, err := h.services.GetUser(message)
		if err != nil {
			if !errors.Is(err, service.ErrUserNotFound) {
				return nil, err
			}

			_, createErr := h.services.CreateUser(message)
			if createErr != nil {
				return nil, fmt.Errorf("failed to create user: %w", createErr)
			}
		}

		startMessage, err := h.services.GetStartMessage(message)
		if err != nil {
			return nil, err
		}

		return &models.MessageResponse{Message: startMessage, Poll: nil}, nil

	case "/menu", "Back to Menu 🏠":
		message, err := h.services.GetMainMenuMessage(message.Chat.ID)
		if err != nil {
			return nil, err
		}
		return &models.MessageResponse{Message: message, Poll: nil}, nil

	case "/topics", "Topics 📖":
		message, err := h.services.GetTopicMenuMessage(message.Chat.ID)
		if err != nil {
			return nil, err
		}
		return &models.MessageResponse{Message: message, Poll: nil}, nil
	case "/progress", "Your progress 🎯":
		message, err := h.services.GetUserProgress(message.Chat.ID)
		if err != nil {
			return nil, err
		}
		return &models.MessageResponse{Message: message, Poll: nil}, nil

	default:
		pollConfig, err := h.services.GetExercisePoll(message.Chat.ID, message.Text)
		if err != nil {
			return nil, err
		}
		return &models.MessageResponse{Poll: pollConfig, Message: nil}, nil
	}
}

func (h *Handler) HandlePollAnswer(pollAnswer *tgbotapi.PollAnswer) error {
	err := h.services.UpdateTopicProgress(pollAnswer.User.ID, pollAnswer.OptionIDs[0])
	if err != nil {
		return fmt.Errorf("failed to update topic progress: %w", err)
	}
	return nil
}
