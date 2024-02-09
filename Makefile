
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
	go build -o cli cmd/cli/main.go

.PHONY: build-api
build-api:
	go build -o api cmd/api/main.go
