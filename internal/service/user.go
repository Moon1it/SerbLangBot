package service

import (
	"fmt"
	"strings"
	"time"

	"github.com/Moon1it/SerbLangBot/internal/models"
	"github.com/Moon1it/SerbLangBot/internal/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	userRepo  repository.User
	topicRepo repository.Topic
}

func InitUserService(userRepo repository.User, topicRepo repository.Topic) *UserService {
	return &UserService{
		userRepo:  userRepo,
		topicRepo: topicRepo,
	}
}

func (u *UserService) GetUser(message *tgbotapi.Message) (*models.User, error) {
	existingUser, err := u.userRepo.GetByChatID(message.Chat.ID)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, ErrUserNotFound
		}
		return nil, fmt.Errorf("failed to get user: %v", err)
	}
	return existingUser, nil
}

func (u *UserService) CreateUser(message *tgbotapi.Message) (*models.User, error) {
	// Get the total count of topics
	count, err := u.topicRepo.GetTopicsCount()
	if err != nil {
		return nil, fmt.Errorf("failed to get topics count: %w", err)
	}

	// Initialize topic progress for all topics
	newProgressByTopics := make([]models.TopicStats, count)
	for i := int64(0); i < count; i++ {
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
	if err := u.userRepo.Create(&newUser); err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return &newUser, nil
}

// This function retrieves user statistics from the database,
// generates a message with this statistics, adds a keyboard,
// and sends the message through a Telegram bot.
func (u *UserService) GetUserProgress(chatID int64) (*tgbotapi.MessageConfig, error) {
	stats, err := u.userRepo.GetStatsByChatID(chatID)
	if err != nil {
		return nil, fmt.Errorf("failed to get UserStats: %w", err)
	}

	message, err := u.generateProgressMessage(stats)
	if err != nil {
		return nil, fmt.Errorf("failed to generate StatsMessage: %w", err)
	}

	msg := tgbotapi.NewMessage(chatID, message)
	msg.ParseMode = tgbotapi.ModeMarkdown

	return &msg, nil
}

// This function generates a message summarizing user progress on
// different topics, including total and successful problems solved.
func (u *UserService) generateProgressMessage(stats *models.UserStats) (string, error) {
	allTopics, err := u.topicRepo.GetAllTopics()
	if err != nil {
		return "", fmt.Errorf("failed to get topics: %w", err)
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
func (u *UserService) UpdateTopicProgress(chatID int64, answer int) error {
	user, err := u.userRepo.GetByChatID(chatID)
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

	err = u.userRepo.UpdateTopicProgress(chatID, int64(topicID), &topicProgress)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}
