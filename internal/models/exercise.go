package models

type Exercise struct {
	ID       string   `bson:"_id"`      // Unique identifier for the exercise
	TopicID  int      `bson:"topicID"`  // Numeric identifier for the topic associated with the exercise
	Question string   `bson:"question"` // The text of the exercise question
	Variants []string `bson:"variants"` // List of answer variants or choices
	Answer   int      `bson:"answer"`   // Index of the correct answer in the 'Variants' slice
}
