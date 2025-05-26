package main

import (
	"bufio"
	"fmt"
	"os"
)

func Error(err error) error {
	if err != nil {
		return err
	}
	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	Error(err)
	fmt.Println(input)
	/*
		 	err := godotenv.Load()
			if err != nil {
				log.Fatal("Error ao carregar variavel!")
			}

			c := openai.NewOpenAIClient(os.Getenv("OPENAI_API_KEY"))
			ctx := context.Background()

			resp, err := chat.SimpleOpenAIPrompt("Bom dia", ctx, c.Client)
			if err != nil {
				fmt.Errorf("Error %v", err)
			}


			fmt.Println(resp)
	*/
}
