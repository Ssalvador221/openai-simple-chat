package openai

import (
	openai "github.com/sashabaranov/go-openai"
)

type Client struct {
	*openai.Client
}

func NewOpenAIClient(apiKey string) *Client {
	client := openai.NewClient(apiKey)

	return &Client{
		client,
	}
}
