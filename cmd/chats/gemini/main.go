package main

import (
	"bufio"
	"chat-ia-go/internal/chat"
	"chat-ia-go/internal/gemini"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/glamour"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	reader := bufio.NewReader(os.Stdin)

	c := gemini.NewGenAIClient(os.Getenv("GEMINI_API_KEY"), ctx)
	defer c.Close()

	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error", err)
		}

		go func() {
			resp, err := chat.SimpleGenAIPrompt(input, ctx, c.Client)
			if err != nil {
				log.Fatal(err.Error())
			}
			out, err := glamour.Render(resp, "dark")
			if err != nil {
				log.Fatal(err.Error())
			}

			fmt.Println(out)
		}()

	}

}
