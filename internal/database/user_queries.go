package database

import (
	"context"
	"strconv"

	"github.com/Moon1it/SerbLangBot/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CreateUser creates a new user in the database.
func CreateUser(db *mongo.Database, newUser models.User) error {
	users := db.Collection("Users")
	// Insert the new user into the database
	_, err := users.InsertOne(context.TODO(), newUser)
	return err
}

func GetUser(db *mongo.Database, chatID int64) (models.User, error) {
	// Get the Users collection
	users := db.Collection("Users")

	// Check if a user with the specified chatID already exists
	filter := bson.M{"chatID": chatID}
	var existingUser models.User
	err := users.FindOne(context.TODO(), filter).Decode(&existingUser)
	if err != nil {
		return models.User{}, err
	}

	return existingUser, nil
}

func GetUserStats(db *mongo.Database, chatID int64) (models.UserStats, error) {
	// Get the Users collection
	users := db.Collection("Users")

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

func UpdateUserCurrentExercise(db *mongo.Database, chatID int64, currentExercise models.Exercise) error {
	// Get the Users collection
	users := db.Collection("Users")
	// Check if a user with the specified chatID already exists
	filter := bson.M{"chatID": chatID}
	update := bson.M{"$set": bson.M{"currentExercise": currentExercise}}

	_, err := users.UpdateOne(context.TODO(), filter, update)
	return err
}

func UpdateUserTopicProgress(db *mongo.Database, chatID int64, topicID int64, topicProgress models.TopicProgress) error {
	// Get the Users collection
	users := db.Collection("Users")
	// Check if a user with the specified chatID already exists
	filter := bson.M{"chatID": chatID}
	update := bson.M{"$set": bson.M{"stats.topicProgress." + strconv.FormatInt(topicID, 10): topicProgress}}

	_, err := users.UpdateOne(context.TODO(), filter, update)
	return err
}
