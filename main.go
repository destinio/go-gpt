package main

import (
	"fmt"
	"os"
	"strings"

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

	args := os.Args
	if len(args) < 2 {
		fmt.Println("ask me something bro!")
		panic("ask me something bro!")
	}

	question := strings.Join(args[1:], " ")

	ai.Init(apiKey, question)

}
