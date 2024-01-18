cli-run:
	go run cmd/clize/main.go

test:
	go test ./...

.PHONY: cli-run