package utils

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Environment int

func getEnvVar(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
	}

	return os.Getenv(key)
}

const (
	BotToken  Environment = iota
	DallE     Environment = iota
	PixBayKey Environment = iota
)

var tokenValues = map[Environment]string{
	BotToken:  getEnvVar("BOT_TOKEN"),
	DallE:     getEnvVar("DallE_KEY"),
	PixBayKey: getEnvVar("PIXBAY_KEY"),
}

func (envName Environment) GetValue() string {
	return tokenValues[envName]
}
