BINARY_NAME=myapp

MAIN=./cmd/chat/main.go

all: build

build:
	go build -o $(BINARY_NAME) $(MAIN)

run:
	go run $(MAIN)

clean:
	rm -f $(BINARY_NAME)
