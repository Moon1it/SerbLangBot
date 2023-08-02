package service

import (
	"fmt"
	"strings"

	"github.com/Moon1it/SerbLangBot/internal/database"
	"github.com/Moon1it/SerbLangBot/internal/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUserStats(db *mongo.Database, chatID int64) (tgbotapi.MessageConfig, error) {
	stats, err := database.GetUserStats(db, chatID)
	if err != nil {
		return tgbotapi.MessageConfig{}, fmt.Errorf("failed to get UserStats: %w", err)
	}

	message, err := generateStatsMessage(db, stats)
	if err != nil {
		return tgbotapi.MessageConfig{}, fmt.Errorf("failed to generate StatsMessage: %w", err)
	}
	msg := tgbotapi.NewMessage(chatID, message)
	GetMainKeyboard(&msg)
	msg.ParseMode = tgbotapi.ModeMarkdown

	return msg, nil
}

func generateStatsMessage(db *mongo.Database, stats models.UserStats) (string, error) {
	message := "Your Progress:\n\n"

	topics, err := database.GetAllTopics(db)
	if err != nil {
		return "", err
	}

	for topicID, topicProgress := range stats.ProgressByTopics {
		topicName := getTopicNameByID(topics, topicID)
		topicMessage := fmt.Sprintf("*%s*:\n", topicName)
		topicMessage += fmt.Sprintf("Total solved: %d\n", topicProgress.AllSolved)
		topicMessage += fmt.Sprintf("Successfully solved: %d\n\n", topicProgress.SuccessfulSolved)

		message += topicMessage
	}

	return strings.TrimSpace(message), nil
}

func getTopicNameByID(topics []models.Topic, topicID int64) string {
	for _, topic := range topics {
		if int64(topic.TopicID) == topicID {
			return topic.Name
		}
	}
	return "" // Return an empty string if topic with given ID is not found
}

func UpdateTopicProgress(db *mongo.Database, chatID int64, answer int) error {
	user, err := database.GetUser(db, chatID)
	if err != nil {
		return fmt.Errorf("failed to get user: %w", err)
	}

	topicID := int64(user.CurrentExercise.TopicID)
	topicProgress, ok := user.Stats.ProgressByTopics[topicID]
	if !ok {
		// If the topicProgress entry doesn't exist, create a new one
		topicProgress = models.TopicProgress{}
	}

	// Update the topicProgress based on the answer
	if user.CurrentExercise.Answer == answer {
		topicProgress.SuccessfulSolved++
	}

	topicProgress.AllSolved++

	// Update the map with the modified topicProgress
	user.Stats.ProgressByTopics[topicID] = topicProgress

	err = database.UpdateUserTopicProgress(db, chatID, topicID, topicProgress)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}
