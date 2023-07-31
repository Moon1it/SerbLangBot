package database

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	"github.com/Moon1it/serb-lang-bot/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// func CreateExercise(db *mongo.Database, question string, answer int, topicId int, variants []string) error {
// 	// Create a new exercise
// 	newExercise := models.Exercise{
// 		ID:       uuid.New().String(),
// 		TopicId:  topicId,
// 		Question: question,
// 		Variants: variants,
// 		Answer:   answer,
// 	}
//
// 	// Insert the exercise into the collection
// 	exercises := db.Collection("Exercises")
// 	_, err := exercises.InsertOne(context.TODO(), newExercise)
// 	if err != nil {
// 		return err
// 	}
//
// 	return nil
// }

// Example usage:
// err := database.CreateExercise(a.db, "questions", "new question", "answer", []string{
//     "Ko?", "Sta?", "Kako?", "Zasto?",
// })

func GetRandomExercise(db *mongo.Database, topicId int) (*models.Exercise, error) {
	// Get the exercises collection
	exercises := db.Collection("Exercises")

	// Create a filter to search for exercises by topic
	filter := bson.M{"topicId": topicId}

	// Prepare parameters to find a random document
	opts := options.Find().SetLimit(1)
	count, err := exercises.CountDocuments(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	opts.SetSkip(rand.Int63n(count))

	// Find a random exercise in the collection based on the filter
	cursor, err := exercises.Find(context.TODO(), filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	// Get the first document from the cursor
	var exercise models.Exercise
	if cursor.Next(context.TODO()) {
		err := cursor.Decode(&exercise)
		if err != nil {
			return nil, err
		}
	} else {
		log.Println("Exercise not found")
		return nil, nil
	}

	return &exercise, nil
}

func GetExerciseByTopicName(db *mongo.Database, topicName string) (*models.Exercise, error) {
	// Get the Exercises collection
	exercises := db.Collection("Exercises")
	pipeline := []bson.M{
		{
			"$lookup": bson.M{
				"from":         "Topics",
				"localField":   "topicId",
				"foreignField": "topicId",
				"as":           "Topic",
			},
		},
		{
			"$match": bson.M{
				"Topic.name": topicName,
			},
		},
		{
			"$sample": bson.M{"size": 1},
		},
		{
			"$project": bson.M{
				"_id":   0,
				"Topic": 0,
			},
		},
	}

	// Perform data aggregation
	cursor, err := exercises.Aggregate(context.Background(), pipeline)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Extract the query results
	var result []*models.Exercise
	if err := cursor.All(context.Background(), &result); err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Close the cursor
	cursor.Close(context.Background())

	// Check if there are any results
	if len(result) == 0 {
		return nil, fmt.Errorf("No exercises found for topic: %s", topicName)
	}

	// Return a random exercise
	return result[0], nil
}

func GetAnswerByQuestion(db *mongo.Database, question string) (*models.Exercise, error) {
	// Get the exercises collection
	exercises := db.Collection("Exercises")

	// Create a filter to search for exercises by question
	filter := bson.M{"question": question}

	// Find a random exercise in the collection based on the filter
	var exercise models.Exercise
	err := exercises.FindOne(context.TODO(), filter).Decode(&exercise)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("exercise not found")
		}
		return nil, fmt.Errorf("failed to get exercise: %w", err)
	}

	return &exercise, nil
}
