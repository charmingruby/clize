BINARY_PATH = ./bin

.PHONY: start
start:
	go run cmd/api/main.go

.PHONY: test
test:
	go test ./...

.PHONY: build
build: build-api build-cli


.PHONY: build-cli
build-cli: 
	go build -o ${BINARY_PATH}/cli cmd/cli/main.go

.PHONY: build-api
build-api:
	go build -o ${BINARY_PATH}/api cmd/api/main.go

.PHONY: run
run:
	./bin/main

