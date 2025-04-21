package openai

import (
	openai "github.com/sashabaranov/go-openai"
)

type Client struct {
	*openai.Client
}

func NewClient(apiKey string) *Client {
	client := openai.NewClient(apiKey)

	return &Client{
		client,
	}
}
