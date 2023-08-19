package services

import (
	"fmt"
	"sort"
	"strings"

	"github.com/Moon1it/SerbLangBot/internal/database"
	"github.com/Moon1it/SerbLangBot/internal/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// This function retrieves user statistics from the database,
// generates a message with this statistics, adds a keyboard,
// and sends the message through a Telegram bot.
func GetUserProgress(chatID int64) (tgbotapi.MessageConfig, error) {
	stats, err := database.GetUserStats(chatID)
	if err != nil {
		return tgbotapi.MessageConfig{}, fmt.Errorf("failed to get UserStats: %w", err)
	}

	message, err := generateStatsMessage(stats)
	if err != nil {
		return tgbotapi.MessageConfig{}, fmt.Errorf("failed to generate StatsMessage: %w", err)
	}
	msg := tgbotapi.NewMessage(chatID, message)
	GetMainKeyboard(&msg)
	msg.ParseMode = tgbotapi.ModeMarkdown

	return msg, nil
}

// This function generates a message summarizing user progress on
// different topics, including total and successful problems solved.
func generateStatsMessage(stats models.UserStats) (string, error) {
	message := "Your Progress:\n\n"

	topics, err := database.GetAllTopics()
	if err != nil {
		return "", err
	}

	// Sort the topics by name in alphabetical order
	sort.SliceStable(topics, func(i, j int) bool {
		return topics[i].Name < topics[j].Name
	})

	for topicID, topicProgress := range stats.ProgressByTopics {
		topicName := getTopicNameByID(topics, topicID)
		if topicName == "" {
			continue // Skip if the topic is not found
		}

		topicMessage := fmt.Sprintf("*%s*:\n", topicName)
		topicMessage += fmt.Sprintf("Total solved: %d\n", topicProgress.AllSolved)
		topicMessage += fmt.Sprintf("Successfully solved: %d\n\n", topicProgress.SuccessfulSolved)

		message += topicMessage
	}

	return strings.TrimSpace(message), nil
}

// This function retrieves the name of a topic based on its ID from a list
// of topics. If the topic ID is not found, it returns an empty string.
func getTopicNameByID(topics []models.Topic, topicID int64) string {
	for _, topic := range topics {
		if int64(topic.TopicID) == topicID {
			return topic.Name
		}
	}
	return "" // Return an empty string if topic with given ID is not found
}

// This function updates the progress of a user's active exercise within
// a specific topic, tracking successful and total solutions, and stores
// the updated information in the database.
func UpdateTopicProgress(chatID int64, answer int) error {
	user, err := database.GetUser(chatID)
	if err != nil {
		return fmt.Errorf("failed to get user: %w", err)
	}

	topicID := int64(user.ActiveExercise.TopicID)
	topicProgress, ok := user.Stats.ProgressByTopics[topicID]
	if !ok {
		// If the topicProgress entry doesn't exist, create a new one
		topicProgress = models.TopicStats{}
	}

	// Update the topicProgress based on the answer
	if user.ActiveExercise.Answer == answer {
		topicProgress.SuccessfulSolved++
	}

	topicProgress.AllSolved++

	// Update the map with the modified topicProgress
	user.Stats.ProgressByTopics[topicID] = topicProgress

	err = database.UpdateUserTopicProgress(chatID, topicID, topicProgress)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}
