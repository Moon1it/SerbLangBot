package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/Moon1it/SerbLangBot/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MessageRepo struct {
	Collection *mongo.Collection
}

func InitMessageRepo(db *mongo.Client) *MessageRepo {
	dbName := os.Getenv("DATABASE")
	messageCollection := db.Database(dbName).Collection("Messages")
	return &MessageRepo{Collection: messageCollection}
}

// GetServiceMessage gets a service message by its name from the database.
// Returns the found service message or an empty message if not found.
func (m *MessageRepo) GetServiceMessage(name string) (*models.ServiceMessage, error) {
	// Define a filter to find service messages by name
	filter := bson.M{"name": name}

	// Execute the query and retrieve the service message
	var message models.ServiceMessage
	err := m.Collection.FindOne(context.TODO(), filter).Decode(&message)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// Handle case when no matching service message is found
			return nil, fmt.Errorf("failed to get service message: %w", err)
		}
		return nil, err
	}
	return &message, nil
}
