.PHONY: build run test clean

build:
	go build -o bin/server cmd/server/main.go

run:
	go run cmd/server/main.go

test:
	go test ./...

clean:
	rm -rf bin/

deps:
	go mod download
	go mod tidy
