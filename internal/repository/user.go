package repository

import (
	"context"
	"os"
	"strconv"

	"github.com/Moon1it/SerbLangBot/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepo struct {
	Collection *mongo.Collection
}

func InitUserRepo(db *mongo.Client) *UserRepo {
	dbName := os.Getenv("DATABASE")
	userCollection := db.Database(dbName).Collection("Users")
	return &UserRepo{Collection: userCollection}
}

// CreateUser creates a new user in the database.
func (u *UserRepo) Create(newUser *models.User) error {
	// Insert the new user into the database
	_, err := u.Collection.InsertOne(context.TODO(), *newUser)
	return err
}

// GetUser finds the user in the database.
func (u *UserRepo) GetByChatID(chatID int64) (*models.User, error) {
	// Check if a user with the specified chatID already exists
	filter := bson.M{"chatID": chatID}
	var existingUser models.User
	err := u.Collection.FindOne(context.TODO(), filter).Decode(&existingUser)
	if err != nil {
		return nil, err
	}
	return &existingUser, nil
}

// GetUserStats finds the user in the database and return only Stats.
func (u *UserRepo) GetStatsByChatID(chatID int64) (*models.UserStats, error) {

	// Check if a user with the specified chatID already exists
	filter := bson.M{"chatID": chatID}

	// Specify that we only need the "stats" field
	projection := bson.M{"stats": 1}

	var existingUser models.User
	err := u.Collection.FindOne(context.TODO(), filter, options.FindOne().SetProjection(projection)).Decode(&existingUser)
	if err != nil {
		return nil, err
	}

	return &existingUser.Stats, nil
}

// UpdateUserActiveExercise updates the active exercise for a user in the database.
func (u *UserRepo) UpdateActiveExercise(chatID int64, activeExercise *models.Exercise) error {

	// Check if a user with the specified chatID already exists
	filter := bson.M{"chatID": chatID}
	update := bson.M{"$set": bson.M{"activeExercise": *activeExercise}}
	_, err := u.Collection.UpdateOne(context.TODO(), filter, update)
	return err
}

// UpdateUserTopicProgress updates the topic progress for a user in the database.
func (u *UserRepo) UpdateTopicProgress(chatID int64, topicID int64, topicProgress *models.TopicStats) error {

	// Check if a user with the specified chatID already exists
	filter := bson.M{"chatID": chatID}
	update := bson.M{"$set": bson.M{"stats.progressByTopics." + strconv.FormatInt(topicID, 10): *topicProgress}}

	_, err := u.Collection.UpdateOne(context.TODO(), filter, update)
	return err
}
