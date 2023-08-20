package models

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type ServiceMessage struct {
	ID   string `bson:"_id"`
	Name string `bson:"name"`
	Text string `bson:"text"`
}

type MessageResponse struct {
	Poll    tgbotapi.SendPollConfig
	Message tgbotapi.MessageConfig
}
