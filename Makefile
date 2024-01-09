include .env
export

swag:
	swag init -g cmd/main.go

deps:
	go mod download

build: deps
	#go build -o cmd/$(APP_IMAGE_NAME) cmd/main.go
	go build -o main cmd/main.go

run: build
	#./$(APP_IMAGE_NAME)
	./main

test: deps
	go generate ./...
	go test -v ./...