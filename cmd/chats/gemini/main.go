package main

import (
	"chat-ia-go/internal/chat"
	"chat-ia-go/internal/gemini"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()

	c := gemini.NewGenAIClient(os.Getenv("GEMINI_API_KEY"), ctx)
	defer c.Close()

	resp, err := chat.SimpleGenAIPrompt("Me conte uma historia engraçada sobre você ou do mundo do conto de fadas!", ctx, c.Client)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(resp)
}
