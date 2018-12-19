GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOMOD=$(GOCMD) mod

DIST=./dist
BINARY_NAME=easycsr
BINARY_NAME_WINDOWS=$(BINARY_NAME).exe

.PHONY: all restore build-prepare build build-unix build-windows test clean

all: restore test build

restore:
	$(GOMOD) download

build-prepare:
	mkdir -p $(DIST)
	
build: build-unix build-windows

build-unix: build-prepare
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(DIST)/$(BINARY_NAME) -v ./main.go

build-windows: build-prepare
	GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(DIST)/$(BINARY_NAME_WINDOWS) -v ./main.go

test:
	$(GOTEST) -v -cover ./...

clean:
	$(GOCLEAN)
	rm -rf $(DIST)
