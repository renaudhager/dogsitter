NAME       := dogsitter
VERSION    :=$(shell git describe --abbrev=0 --tags --exact-match 2>/dev/null || git rev-parse --short HEAD)
LDFLAGS    := -w -extldflags "-static" -X 'main.version=$(VERSION)'

ifndef GOBIN
GOBIN = $(GOPATH)/bin
endif

.PHONY: setup
setup:
	go get -u -v golang.org/x/lint/golint
	go get -u -v github.com/mitchellh/gox
	go get -u -v github.com/golang/dep/cmd/dep

.PHONY: deps
deps:
	dep ensure -v

.PHONY: lint
lint:
	golint -set_exit_status .
	go vet ./...

.PHONY: test
test:
	go test -coverprofile=coverage.out -v ./...

.PHONY: build
build:
	mkdir -p dist; rm -rf dist/*
	CGO_ENABLED=0 gox -osarch "linux/386 linux/amd64 darwin/amd64" -ldflags "$(LDFLAGS)" -output dist/$(NAME)_{{.OS}}_{{.Arch}}
	strip dist/*_linux_amd64 dist/*_linux_386
