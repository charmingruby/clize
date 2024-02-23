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
build:
	go build -o api cmd/api/main.go

#####################
# COMMANDS         #
#####################
.PHONY: c
c: 
	go run cmd/cli/main.go $(cmd)