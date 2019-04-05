.PHONY: all
all: build

get:
	go get -v ./...

build:
	go build -o $(GOBIN)/dogsitter dogsitter/main.go
