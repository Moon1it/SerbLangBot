package service

import (
	"errors"

	"github.com/Moon1it/SerbLangBot/internal/models"
	"github.com/Moon1it/SerbLangBot/internal/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

var ErrUserNotFound = errors.New("user not found")

type User interface {
	CreateUser(message *tgbotapi.Message) (*models.User, error)
	GetUser(message *tgbotapi.Message) (*models.User, error)
	UpdateTopicProgress(chatID int64, answer int) error
	GetUserProgress(chatID int64) (*tgbotapi.MessageConfig, error)
}

type Navigation interface {
	GetStartMessage(message *tgbotapi.Message) (*tgbotapi.MessageConfig, error)
	GetMainMenuMessage(chatID int64) (*tgbotapi.MessageConfig, error)
	GetTopicMenuMessage(chatID int64) (*tgbotapi.MessageConfig, error)
}

type Exercise interface {
	GetExercisePoll(chatID int64, exerciseType string) (*tgbotapi.SendPollConfig, error)
}

type Service struct {
	User
	Navigation
	Exercise
}

func InitService(repos *repository.Repository) *Service {
	return &Service{
		User:       InitUserService(repos.User, repos.Topic),
		Navigation: InitNavigationService(repos.User, repos.Message, repos.Topic),
		Exercise:   InitExerciseService(repos.User, repos.Exercise),
	}
}
