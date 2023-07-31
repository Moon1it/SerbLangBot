package config

import (
	"os"
)

type Config struct {
	MongoString string

	// DBHost     string
	// DBPort     string
	// DBUsername string
	DBName string
	// DBAuth     string
	// DBPassword string
}

func Init() (*Config, error) {
	var cfg Config = Config{
		MongoString: os.Getenv("MongoString"),

		// DBHost:     os.Getenv("DB_HOST"),
		// DBPort:     os.Getenv("DB_PORT"),
		// DBUsername: os.Getenv("DB_USERNAME"),
		DBName: os.Getenv("DB_NAME"),
		// DBPassword: os.Getenv("DB_PASSWORD"),
	}
	return &cfg, nil
}
