package ai

import (
	"bytes"
	"context"
	"fmt"

	openai "github.com/sashabaranov/go-openai"
	markdown "github.com/yuin/goldmark"
	"github.com/yuin/goldmark/util"
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

	// result := markdown.Render(resp.Choices[0].Message.Content, 80, 6)
	var buf bytes.Buffer

	if err := markdown.Convert(util.StringToReadOnlyBytes(resp.Choices[0].Message.Content), &buf); err != nil {
		panic(err)
	}

	fmt.Println(buf.String())
}
