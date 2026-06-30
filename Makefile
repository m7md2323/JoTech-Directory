.PHONY: all build run test lint fmt clean docker-build

APP_NAME=server

all: fmt lint test build

run:
	templ generate
	go run ./cmd/server/main.go

build:
	templ generate
	go build -o $(APP_NAME) ./cmd/server

test:
	go test -v ./...

lint:
	templ fmt .
	golangci-lint run

fmt:
	templ fmt .
	go fmt ./...

clean:
	rm -f $(APP_NAME)

docker-build:
	docker build -t jordan-tech-companies .
