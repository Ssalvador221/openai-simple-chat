package main

import (
	"chat-ia-go/internal/chat"
	"chat-ia-go/internal/openai"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error ao carregar variavel!")
	}

	c := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	ctx := context.Background()

	resp, err := chat.SimplePrompt("Bom dia", ctx, c.Client)
	if err != nil {
		fmt.Errorf("Error %v", err)
	}

	fmt.Println(resp)
}
