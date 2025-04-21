package gemini

import (
	"context"
	"log"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type Client struct {
	*genai.Client
}

func NewGenAIClient(env string, ctx context.Context) *Client {
	c, err := genai.NewClient(ctx, option.WithAPIKey(env))
	if err != nil {
		log.Fatal(err)
	}

	return &Client{
		c,
	}
}
