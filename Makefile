start:
	go run cmd/api/main.go

test:
	go test ./...

build:
	go build -o ./bin  cmd/api/main.go

run:
	./bin/main

.PHONY: start run test build
