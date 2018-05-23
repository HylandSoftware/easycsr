GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

DIST=./dist
BINARY_NAME=easycsr
BINARY_NAME_WINDOWS=$(BINARY_NAME).exe

.PHONY: all build-prepare build build-unix build-windows test clean

all: test build

build-prepare:
	mkdir -p $(DIST)
	
build: build-unix build-windows

build-unix: build-prepare
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(DIST)/$(BINARY_NAME) -v ./cmd/easycsr

build-windows: build-prepare
	GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(DIST)/$(BINARY_NAME_WINDOWS) -v ./cmd/easycsr

test:
	$(GOTEST) -v -cover ./...

clean:
	$(GOCLEAN)
	rm -rf $(DIST)
