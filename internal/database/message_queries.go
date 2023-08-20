package database

import (
	"context"

	"github.com/Moon1it/SerbLangBot/internal/models"
	"github.com/Moon1it/SerbLangBot/pkg/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetServiceMessage(name string) (models.ServiceMessage, error) {
	messages := database.GetCollection("Messages")

	// Define a filter to find service messages by name
	filter := bson.M{"name": name}

	// Execute the query and retrieve the service message
	var message models.ServiceMessage
	err := messages.FindOne(context.TODO(), filter).Decode(&message)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// Handle case when no matching service message is found
			return models.ServiceMessage{}, nil
		}
		return models.ServiceMessage{}, err
	}

	// Continue with any additional processing or return the service message
	return message, nil
}
