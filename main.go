package main

import (
	"fmt"
	"os"

	"github.com/destinio/go-gpt/ai"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		fmt.Println("OPENAI_API_KEY is not set")
		panic("OPENAI_API_KEY is not set")
	}

	ai.Init(apiKey)

}
