package services

import (
	"fmt"
	"strings"
	"time"

	"github.com/Moon1it/SerbLangBot/internal/database"
	"github.com/Moon1it/SerbLangBot/internal/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/uuid"
)

func CreateUser(message *tgbotapi.Message) (models.User, error) {
	// Get the total count of topics
	count, err := database.GetTopicsCount()
	if err != nil {
		return models.User{}, fmt.Errorf("failed to get topics count: %w", err)
	}

	// Initialize topic progress for all topics
	newProgressByTopics := make([]models.TopicStats, count)
	var i int64
	for i = 0; i < count; i++ {
		newProgressByTopics[i] = models.TopicStats{
			AllSolved:        0,
			SuccessfulSolved: 0,
		}
	}

	stats := models.UserStats{
		ProgressByTopics: newProgressByTopics,
	}

	newUser := models.User{
		ID:             uuid.New().String(),
		Name:           message.From.UserName,
		ChatID:         message.Chat.ID,
		ActiveExercise: models.Exercise{},
		Stats:          stats,
		RegisteredAt:   time.Now(),
	}

	// Create the new user in the database
	if err := database.CreateUser(newUser); err != nil {
		return models.User{}, fmt.Errorf("failed to create user: %w", err)
	}

	return newUser, nil
}

// This function retrieves user statistics from the database,
// generates a message with this statistics, adds a keyboard,
// and sends the message through a Telegram bot.
func GetUserProgress(chatID int64) (tgbotapi.MessageConfig, error) {
	stats, err := database.GetUserStats(chatID)
	if err != nil {
		return tgbotapi.MessageConfig{}, fmt.Errorf("failed to get UserStats: %w", err)
	}

	message, err := generateProgressMessage(stats)
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
func generateProgressMessage(stats models.UserStats) (string, error) {
	allTopics, err := database.GetAllTopics()
	if err != nil {
		return "", err
	}

	message := "Your Progress:\n\n"

	// Iterate through the sorted topic IDs and generate the message
	for index, topic := range stats.ProgressByTopics {

		topicMessage := fmt.Sprintf("*%s*:\n", allTopics[index].Name)
		topicMessage += fmt.Sprintf("Total solved: %d\n", topic.AllSolved)
		topicMessage += fmt.Sprintf("Successfully solved: %d\n\n", topic.SuccessfulSolved)

		message += topicMessage
	}

	return strings.TrimSpace(message), nil
}

// This function updates the progress of a user's active exercise within
// a specific topic, tracking successful and total solutions, and stores
// the updated information in the database.
func UpdateTopicProgress(chatID int64, answer int) error {
	user, err := database.GetUser(chatID)
	if err != nil {
		return fmt.Errorf("failed to get user: %w", err)
	}

	topicID := user.ActiveExercise.TopicID
	topicProgress := user.Stats.ProgressByTopics[topicID]

	// Update the topicProgress based on the answer
	if user.ActiveExercise.Answer == answer {
		topicProgress.SuccessfulSolved++
	}

	topicProgress.AllSolved++

	// Update the array with the modified topicProgress
	user.Stats.ProgressByTopics[topicID] = topicProgress

	err = database.UpdateUserTopicProgress(chatID, int64(topicID), topicProgress)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}
