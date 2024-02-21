#####################
# SERVER           #
#####################
.PHONY: server
server:
	go run cmd/api/main.go

#####################
# TESTS            #
#####################
.PHONY: test
test:
	go test ./...

#####################
# BUILD            #
#####################
.PHONY: build
build: build-api build-cli

.PHONY: build-cli
build-cli: 
	go build -o cli cmd/cli/main.go

.PHONY: build-api
build-api:
	go build -o api cmd/api/main.go

#####################
# COMMANDS         #
#####################
.PHONY: c
c: 
	go run cmd/cli/main.go $(cmd)