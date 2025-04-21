package chat

import (
	"context"
	"fmt"
	"log"

	"github.com/google/generative-ai-go/genai"
	openai "github.com/sashabaranov/go-openai"
)

func SimpleOpenAIPrompt(prompt string, ctx context.Context, c *openai.Client) (string, error) {
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

func SimpleGenAIPrompt(prompt string, ctx context.Context, c *genai.Client) (string, error) {
	model := c.GenerativeModel("gemini-2.0-flash")

	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		log.Fatal(err.Error())
	}

	text := resp.Candidates[0].Content.Parts[0].(genai.Text)

	return string(text), err
}
