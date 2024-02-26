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
build: build-server build-cli

.PHONY: build-server
build-server:
	go build -o api cmd/api/main.go

.PHONY: build-cli
build-cli:
	go build -o clize cmd/cli/main.go

#####################
# COMMANDS         #
#####################
.PHONY: c
c: 
	go run cmd/cli/main.go $(cmd)