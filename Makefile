NAME       := dogsitter
VERSION    :=$(shell git describe --abbrev=0 --tags --exact-match 2>/dev/null || git rev-parse --short HEAD)
LDFLAGS    := -w -extldflags "-static" -X 'main.version=$(VERSION)'

ifndef GOBIN
GOBIN = $(GOPATH)/bin
endif

.PHONY: setup
setup:
	go get -u -v golang.org/x/lint/golint
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

.PHONY: build-docker
build:
	CGO_ENABLED=0 go build -o $(NAME) -ldflags "$(LDFLAGS)" .
	strip $(NAME)
	cp $(NAME) $(NAME)_$(VERSION)
