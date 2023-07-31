package service

import (
	"fmt"

	"github.com/Moon1it/serb-lang-bot/internal/database"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetExerciseByTopic(db *mongo.Database, chatID int64, exerciseType string) (tgbotapi.SendPollConfig, error) {
	exercise, err := database.GetExerciseByTopicName(db, exerciseType)
	if err != nil {
		return tgbotapi.SendPollConfig{}, fmt.Errorf("failed to get exercise by topic name: %w", err)
	}

	msg := tgbotapi.NewPoll(chatID, exercise.Question, exercise.Variants...)
	msg.Type = "quiz"
	msg.IsAnonymous = false
	msg.CorrectOptionID = int64(exercise.Answer) // Set the correct option ID

	err = database.UpdateUserCurrentExercise(db, chatID, *exercise)
	if err != nil {
		return tgbotapi.SendPollConfig{}, fmt.Errorf("failed to update user: %w", err)
	}

	return msg, nil
}
