package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerAddress string
	DbUsername    string
	DbPassword    string
	DbPort        string
	DbHost        string
	DbName        string
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func InitConfiguration() Config {
	return Config{
		ServerAddress: os.Getenv("SERVER_ADDRESS"),
		DbUsername:    os.Getenv("DB_USERNAME"),
		DbPassword:    os.Getenv("DB_PASSWORD"),
		DbPort:        os.Getenv("DB_PORT"),
		DbHost:        os.Getenv("DB_HOST"),
		DbName:        os.Getenv("DB_NAME"),
	}
}
