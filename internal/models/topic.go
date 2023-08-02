package models

type Topic struct {
	ID      string `bson:"_id"`     // Unique identifier for the topic
	TopicID int    `bson:"topicID"` // Numeric identifier for the topic
	Name    string `bson:"name"`    // Name of the topic
}
