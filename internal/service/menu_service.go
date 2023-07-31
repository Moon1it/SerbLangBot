package service

import (
	"fmt"

	"github.com/Moon1it/SerbLangBot/internal/database"
	"github.com/Moon1it/SerbLangBot/internal/models"
	"github.com/Moon1it/SerbLangBot/pkg/clients/keyboard"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetMainKeyboard(msg *tgbotapi.MessageConfig) {
	buttons := []string{"Your progress 🎯", "Topics 📖"}
	msg.ReplyMarkup = keyboard.CreateMenuKeyboard(buttons)
}

func GetStartMessage(db *mongo.Database, chatID int64, userName string) (tgbotapi.MessageConfig, error) {
	// Check if the user already exists in the database
	_, err := database.GetUser(db, chatID)
	if err != nil {
		if err == mongo.ErrNoDocuments { // User does not exist, create a new one
			// Initialize the map for new topic progress
			newTopicProgress := make(map[int64]models.TopicProgress)

			// Get the total count of topics
			count, err := database.GetTopicsCount(db)
			if err != nil {
				return tgbotapi.MessageConfig{}, fmt.Errorf("failed to get topics count: %w", err)
			}

			// Initialize topic progress for all topics
			for i := int64(1); i <= count; i++ {
				newTopicProgress[i] = models.TopicProgress{
					AllSolved:        0,
					SuccessfulSolved: 0,
				}
			}

			stats := models.UserStats{
				TopicProgress: newTopicProgress,
			}

			newUser := models.User{
				ID:              uuid.New().String(),
				Name:            userName,
				ChatID:          chatID,
				CurrentExercise: models.Exercise{},
				Stats:           stats,
			}

			// Create the new user in the database
			if err := database.CreateUser(db, newUser); err != nil {
				return tgbotapi.MessageConfig{}, fmt.Errorf("failed to create user: %w", err)
			}
		} else {
			return tgbotapi.MessageConfig{}, fmt.Errorf("failed to get user: %w", err)
		}
	}

	// The user already exists or has been successfully created
	startMessage := `🌟Hello and welcome to the Serbian Language Learning Bot!🌟

I'm your dedicated language learning companion designed to help you master Serbian. Learning a new language comes with a host of exciting benefits:

📖 Start Learning: Begin your language lessons and practice Serbian at your own pace. Whether you're a complete beginner or looking to enhance your existing knowledge, we've got you covered!

🔄 Daily Practice: Engage in daily exercises to reinforce your language skills. Consistency is the key to language mastery, and our daily practice sessions will keep you on track.

🎯 Vocabulary: Access our comprehensive Serbian vocabulary database. Expand your word bank and express yourself fluently in Serbian.

🎮 Language Games: Have fun and learn through interactive games. Learning doesn't have to be boring – our language games make the process enjoyable and engaging!

📈 Progress & Stats: Track your progress and see how far you've come. Celebrate your achievements and stay motivated as you see yourself making steady progress.

To get started, simply type /menu and explore the various options available in our main menu. Let's embark on this language adventure together! 🚀😊`

	msg := tgbotapi.NewMessage(chatID, startMessage)
	GetMainKeyboard(&msg)
	return msg, nil
}

func GetMainMenu(db *mongo.Database, chatID int64) (tgbotapi.MessageConfig, error) {
	msg := tgbotapi.NewMessage(chatID, "📚 Main Menu 📚\n\nWelcome back! Here are the options available in the main menu:")
	GetMainKeyboard(&msg)
	return msg, nil
}

func GetTopicMenu(db *mongo.Database, chatID int64) (tgbotapi.MessageConfig, error) {
	topicsList, err := database.GetAllTopics(db)
	if err != nil {
		msg := tgbotapi.NewMessage(chatID, "Failed to get topic or command")
		return msg, err
	}
	// Create app slice to store the topic names
	topicNames := make([]string, len(topicsList))

	// Extract the name of each topic and add it to the slice
	for i, topic := range topicsList {
		topicNames[i] = topic.Name
	}

	// Add "Menu" button to the end of the topicNames slice
	topicNames = append(topicNames, "Back to Menu 🏠")

	msg := tgbotapi.NewMessage(chatID, "Choose app topic for practice:")
	msg.ReplyMarkup = keyboard.CreateMenuKeyboard(topicNames)
	return msg, nil
}
