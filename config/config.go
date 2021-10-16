package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var PORT int

func GoDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func InitPort() {
	PORT, _ = strconv.Atoi(GoDotEnvVariable("PORT"))
}
