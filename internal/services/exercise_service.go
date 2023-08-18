package services

import (
	"fmt"

	"github.com/Moon1it/SerbLangBot/internal/database"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func GetExerciseByTopic(chatID int64, exerciseType string) (tgbotapi.SendPollConfig, error) {
	exercise, err := database.GetExerciseByTopicName(exerciseType)
	if err != nil {
		return tgbotapi.SendPollConfig{}, fmt.Errorf("failed to get exercise by topic name: %w", err)
	}

	msg := tgbotapi.NewPoll(chatID, exercise.Question, exercise.Variants...)
	msg.Type = "quiz"
	msg.IsAnonymous = false
	msg.CorrectOptionID = int64(exercise.Answer) // Set the correct option ID

	err = database.UpdateUserActiveExercise(chatID, *exercise)
	if err != nil {
		return tgbotapi.SendPollConfig{}, fmt.Errorf("failed to update user: %w", err)
	}

	return msg, nil
}
