NAME       := dogsitter

ifndef DRONE_TAG
	VERSION :=$(shell git describe --abbrev=0 --tags --exact-match 2>/dev/null || git rev-parse --short HEAD)
else
	VERSION := $(DRONE_TAG)
endif

LDFLAGS    := -w -extldflags "-static" -X 'main.version=$(VERSION)'

ifndef GOBIN
GOBIN = $(GOPATH)/bin
endif

.PHONY: setup
setup:
	GO111MODULE=on go get -u -v golang.org/x/lint/golint
	GO111MODULE=on go get -u -v github.com/mitchellh/gox

.PHONY: deps
deps:
	GO111MODULE=on go mod download

.PHONY: lint
lint:
	GO111MODULE=on golint -set_exit_status .
	GO111MODULE=on go vet ./...

.PHONY: test
test:
	GO111MODULE=on go test -coverprofile=coverage.out -v ./...

.PHONY: build
build:
	mkdir -p dist; rm -rf dist/*
	CGO_ENABLED=0 GO111MODULE=on gox -osarch "linux/386 linux/amd64 darwin/amd64" -ldflags "$(LDFLAGS)" -output dist/$(NAME)_{{.OS}}_{{.Arch}}
	strip dist/*_linux_amd64 dist/*_linux_386

.PHONY: release
release:
	/usr/local/bin/publish_release.bash
