start:
	go run cmd/clize/main.go

test:
	go test ./...

build:
	go build -o ./bin  cmd/clize/main.go

run:
	./bin/main

.PHONY: start run test build
