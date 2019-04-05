.PHONY: all
all: build

setup:
ifndef GOBIN
GOBIN = $(GOPATH)/bin
endif

get:
	go get -v ./...

build: setup
	go build -o $(GOBIN)/dogsitter dogsitter/main.go
