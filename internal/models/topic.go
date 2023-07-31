package models

type Topic struct {
	ID      string `bson:"_id"`
	TopicID int    `bson:"topicId"`
	Name    string `bson:"name"`
}
