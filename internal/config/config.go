package config

import (
	"os"
)

type Config struct {
	MongoString string

	DBName string
	// DBHost     string
	// DBPort     string
	// DBUsername string
	// DBAuth     string
	// DBPassword string
}

func Init() (*Config, error) {
	var cfg Config = Config{
		MongoString: os.Getenv("MONGO_STRING"),

		DBName: os.Getenv("DB_NAME"),
		// DBHost:     os.Getenv("DB_HOST"),
		// DBPort:     os.Getenv("DB_PORT"),
		// DBUsername: os.Getenv("DB_USERNAME"),
		// DBPassword: os.Getenv("DB_PASSWORD"),
	}
	return &cfg, nil
}
