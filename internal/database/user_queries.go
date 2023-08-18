package database

import (
	"context"
	"strconv"

	"github.com/Moon1it/SerbLangBot/internal/models"
	"github.com/Moon1it/SerbLangBot/pkg/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CreateUser creates a new user in the database.
func CreateUser(newUser models.User) error {

	users := database.GetCollection("Users")
	// Insert the new user into the database
	_, err := users.InsertOne(context.TODO(), newUser)
	return err
}

// GetUser finds the user in the database.
func GetUser(chatID int64) (models.User, error) {
	// Get the Users collection
	users := database.GetCollection("Users")

	// Check if a user with the specified chatID already exists
	filter := bson.M{"chatID": chatID}
	var existingUser models.User
	err := users.FindOne(context.TODO(), filter).Decode(&existingUser)
	if err != nil {
		return models.User{}, err
	}

	return existingUser, nil
}

// GetUserStats finds the user in the database and return only Stats.
func GetUserStats(chatID int64) (models.UserStats, error) {
	// Get the Users collection
	users := database.GetCollection("Users")

	// Check if a user with the specified chatID already exists
	filter := bson.M{"chatID": chatID}

	// Specify that we only need the "stats" field
	projection := bson.M{"stats": 1}

	var existingUser models.User
	err := users.FindOne(context.TODO(), filter, options.FindOne().SetProjection(projection)).Decode(&existingUser)
	if err != nil {
		return models.UserStats{}, err
	}

	return existingUser.Stats, nil
}

// UpdateUserActiveExercise updates the active exercise for a user in the database.
func UpdateUserActiveExercise(chatID int64, ActiveExercise models.Exercise) error {
	// Get the Users collection
	users := database.GetCollection("Users")
	// Check if a user with the specified chatID already exists
	filter := bson.M{"chatID": chatID}
	update := bson.M{"$set": bson.M{"activeExercise": ActiveExercise}}
	_, err := users.UpdateOne(context.TODO(), filter, update)
	return err
}

// UpdateUserTopicProgress updates the topic progress for a user in the database.
func UpdateUserTopicProgress(chatID int64, topicID int64, topicProgress models.TopicStats) error {
	// Get the Users collection
	users := database.GetCollection("Users")
	// Check if a user with the specified chatID already exists
	filter := bson.M{"chatID": chatID}
	update := bson.M{"$set": bson.M{"stats.progressByTopics." + strconv.FormatInt(topicID, 10): topicProgress}}

	_, err := users.UpdateOne(context.TODO(), filter, update)
	return err
}
