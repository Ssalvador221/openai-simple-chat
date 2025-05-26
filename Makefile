BINARY_NAME=chat-ai-go

gpt=./cmd/chats/gpt/main.go
gemini=./cmd/chats/gemini/main.go

all: build

build:
	go build -o $(BINARY_NAME) $(MAIN)

run: 
	@echo "Use: make run-gpt ou make run-gemini"

run-gpt:
	go run $(gpt)

run-gemini:
	go run $(gemini)

clean:
	rm -f $(BINARY_NAME)
