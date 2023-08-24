package service

import (
	"fmt"

	keyboard "github.com/Moon1it/SerbLangBot/internal/clients/telegram/keyboard"
	"github.com/Moon1it/SerbLangBot/internal/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type NavigationService struct {
	messageRepo repository.Message
	userRepo    repository.User
	topicRepo   repository.Topic
}

func InitNavigationService(userRepo repository.User, messageRepo repository.Message, topicRepo repository.Topic) *NavigationService {
	return &NavigationService{
		userRepo:    userRepo,
		messageRepo: messageRepo,
		topicRepo:   topicRepo,
	}
}

func (n *NavigationService) GetStartMessage(message *tgbotapi.Message) (*tgbotapi.MessageConfig, error) {
	startMessage, err := n.messageRepo.GetServiceMessage("StartMessage")
	if err != nil {
		return nil, fmt.Errorf("failed to get start message: %w", err)
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, startMessage.Text)
	n.getMainKeyboard(&msg)
	return &msg, nil
}

func (n *NavigationService) GetMainMenuMessage(chatID int64) (*tgbotapi.MessageConfig, error) {
	menuMessage, err := n.messageRepo.GetServiceMessage("MenuMessage")
	if err != nil {
		return nil, fmt.Errorf("failed to get menu message: %w", err)
	}

	msg := tgbotapi.NewMessage(chatID, menuMessage.Text)
	n.getMainKeyboard(&msg)
	return &msg, nil
}

func (n *NavigationService) GetTopicMenuMessage(chatID int64) (*tgbotapi.MessageConfig, error) {
	topicsList, err := n.topicRepo.GetAllTopics()
	if err != nil {
		msg := tgbotapi.NewMessage(chatID, "Failed to get topic or command")
		return &msg, err
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
	return &msg, nil
}

func (n *NavigationService) getMainKeyboard(msg *tgbotapi.MessageConfig) {
	buttons := []string{"Your progress 🎯", "Topics 📖"}
	msg.ReplyMarkup = keyboard.CreateMenuKeyboard(buttons)
}
