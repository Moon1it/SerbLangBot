package service

import (
	"fmt"

	"github.com/Moon1it/SerbLangBot/internal/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type ExerciseService struct {
	userRepo     repository.User
	exerciseRepo repository.Exercise
}

func InitExerciseService(userRepo repository.User, exerciseRepo repository.Exercise) *ExerciseService {
	return &ExerciseService{
		userRepo:     userRepo,
		exerciseRepo: exerciseRepo,
	}
}

func (e *ExerciseService) GetExercisePoll(chatID int64, exerciseType string) (*tgbotapi.SendPollConfig, error) {
	exercise, err := e.exerciseRepo.GetByTopic(exerciseType)
	if err != nil {
		return nil, fmt.Errorf("failed to get exercise by topic name: %w", err)
	}

	msg := tgbotapi.NewPoll(chatID, exercise.Question, exercise.Variants...)
	msg.Type = "quiz"
	msg.IsAnonymous = false
	msg.CorrectOptionID = int64(exercise.Answer) // Set the correct option ID

	err = e.userRepo.UpdateActiveExercise(chatID, exercise)
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	return &msg, nil
}
