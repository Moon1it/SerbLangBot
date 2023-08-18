package database

import (
	"context"
	"errors"
	"fmt"

	"github.com/Moon1it/SerbLangBot/internal/models"
	"github.com/Moon1it/SerbLangBot/pkg/database"
	"go.mongodb.org/mongo-driver/bson"
)

func GetRandomExercise(topicId int) (*models.Exercise, error) {
	exercises := database.GetCollection("Exercises")

	filter := bson.M{"topicId": topicId}

	pipeline := bson.A{
		bson.D{{Key: "$match", Value: filter}},
		bson.D{{Key: "$sample", Value: bson.D{{Key: "size", Value: 1}}}},
	}

	cursor, err := exercises.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return nil, fmt.Errorf("failed to execute aggregation: %w", err)
	}
	defer cursor.Close(context.TODO())

	var exercise models.Exercise
	if cursor.Next(context.TODO()) {
		err := cursor.Decode(&exercise)
		if err != nil {
			return nil, fmt.Errorf("failed to decode exercise: %w", err)
		}
	} else {
		return nil, errors.New("exercise not found")
	}

	return &exercise, nil
}

func GetExerciseByTopicName(topicName string) (*models.Exercise, error) {
	exercises := database.GetCollection("Exercises")

	pipeline := bson.A{
		bson.D{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "Topics"},
				{Key: "localField", Value: "topicID"},
				{Key: "foreignField", Value: "topicID"},
				{Key: "as", Value: "Topic"},
			}},
		},
		bson.D{
			{Key: "$match", Value: bson.M{
				"Topic.name": topicName,
			}},
		},
		bson.D{
			{Key: "$sample", Value: bson.M{"size": 1}},
		},
		bson.D{
			{Key: "$project", Value: bson.M{
				"_id":   0,
				"Topic": 0,
			}},
		},
	}

	cursor, err := exercises.Aggregate(context.Background(), pipeline)
	if err != nil {
		return nil, fmt.Errorf("failed to execute aggregation: %w", err)
	}
	defer cursor.Close(context.Background())

	if !cursor.Next(context.Background()) {
		return nil, fmt.Errorf("no exercises found for topic: %s", topicName)
	}

	var exercise models.Exercise
	if err := cursor.Decode(&exercise); err != nil {
		return nil, fmt.Errorf("failed to decode exercise: %w", err)
	}

	return &exercise, nil
}

func GetAnswerByQuestion(question string) (*models.Exercise, error) {
	exercises := database.GetCollection("Exercises")

	pipeline := bson.A{
		bson.D{
			{Key: "$match", Value: bson.M{"question": question}},
		},
		bson.D{
			{Key: "$sample", Value: bson.M{"size": 1}},
		},
	}

	cursor, err := exercises.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return nil, fmt.Errorf("failed to execute aggregation: %w", err)
	}
	defer cursor.Close(context.TODO())

	if !cursor.Next(context.TODO()) {
		return nil, fmt.Errorf("exercise not found for question: %s", question)
	}

	var exercise models.Exercise
	if err := cursor.Decode(&exercise); err != nil {
		return nil, fmt.Errorf("failed to decode exercise: %w", err)
	}

	return &exercise, nil
}
