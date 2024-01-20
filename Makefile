cli-run:
	go run cmd/clize/main.go

test:
	go test ./...

build:
	go build -o ./bin  cmd/clize/main.go

run:
	./bin/main

.PHONY: cli-run test build