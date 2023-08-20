package services

import (
	"fmt"

	keyboard "github.com/Moon1it/SerbLangBot/internal/clients/telegram/keyboard"
	"github.com/Moon1it/SerbLangBot/internal/database"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetMainKeyboard(msg *tgbotapi.MessageConfig) {
	buttons := []string{"Your progress 🎯", "Topics 📖"}
	msg.ReplyMarkup = keyboard.CreateMenuKeyboard(buttons)
}

func GetStartMessage(message *tgbotapi.Message) (tgbotapi.MessageConfig, error) {
	// Check if the user already exists in the database
	_, err := database.GetUser(message.Chat.ID)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// User does not exist, create a new one
			_, err := CreateUser(message)
			if err != nil {
				return tgbotapi.MessageConfig{}, fmt.Errorf("failed to create user: %w", err)
			}
		} else {
			return tgbotapi.MessageConfig{}, fmt.Errorf("failed to get user: %w", err)
		}
	}

	startMessage, err := database.GetServiceMessage("StartMessage")
	if err != nil {
		return tgbotapi.MessageConfig{}, fmt.Errorf("failed to get start message: %w", err)
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, startMessage.Text)
	GetMainKeyboard(&msg)
	return msg, nil
}

func GetMainMenu(chatID int64) (tgbotapi.MessageConfig, error) {
	menuMessage, err := database.GetServiceMessage("MenuMessage")
	if err != nil {
		return tgbotapi.MessageConfig{}, fmt.Errorf("failed to get menu message: %w", err)
	}

	msg := tgbotapi.NewMessage(chatID, menuMessage.Text)
	GetMainKeyboard(&msg)
	return msg, nil
}

func GetTopicMenu(chatID int64) (tgbotapi.MessageConfig, error) {
	topicsList, err := database.GetAllTopics()
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
