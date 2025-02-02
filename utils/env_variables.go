package utils

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	err := godotenv.Load("server.env")
	if err != nil {
		err = fmt.Errorf("error loading environment variables %w", err)
		log.Fatal(err)
	}
}

func GetValueFromEnv(key string) (string, Status) {
	value, ok := os.LookupEnv(key)
	if !ok {
		return "", Status{Code: 404, Message: "Environment variable not found", Error: errors.New("environment variable not found")}
	}
	return value, Status{Code: 200, Message: "", Error: nil}
}
