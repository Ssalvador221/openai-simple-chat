package chat

import (
	"context"
	"fmt"

	openai "github.com/sashabaranov/go-openai"
)

func SimplePrompt(prompt string, ctx context.Context, c *openai.Client) (string, error) {
	req := openai.ChatCompletionRequest{
		Model: openai.GPT4,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: prompt,
			},
		},
		MaxTokens: 100,
	}

	resp, err := c.CreateChatCompletion(ctx, req)

	if err != nil {
		return "", fmt.Errorf("completition error: %v", err)
	}

	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("no choices returned from completion")
	}

	return resp.Choices[0].Message.Content, nil
}
