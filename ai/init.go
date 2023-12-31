package ai

import (
	"context"
	"fmt"

	"github.com/charmbracelet/glamour"
	openai "github.com/sashabaranov/go-openai"
)

func Init(apiKey string, question string) {

	client := openai.NewClient(apiKey)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: question,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	out, err := glamour.Render(resp.Choices[0].Message.Content, "dark")

	if err != nil {
		panic(err)
	}

	fmt.Print(out)
}
