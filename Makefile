GO ?= go
GOFMT ?= gofmt "-s"
PLATFORM ?=darwin
GOARCH ?=amd64
GOARM ?= 7
GOFILES := $(shell find . -name "*.go" -type f)

all: install

install:
	env GOOS=$(PLATFORM) GOARCH=$(GOARCH) GOARM=$(GOARM) $(GO) build -o ./identixone ./
	mv ./identixone /usr/local/bin/identixone

fmt:
	$(GOFMT) -w $(GOFILES)

test:
	go test -cover ./...