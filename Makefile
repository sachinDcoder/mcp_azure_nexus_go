# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTIDY=$(GOCMD) mod tidy
GOTOOL=$(GOCMD) tool

# Binary name
BINARY_NAME=azure-nexus-mcp-server

all: tidy build

build:
	$(GOBUILD) -o $(BINARY_NAME) .

run: build
	./$(BINARY_NAME)

tidy:
	$(GOTIDY)

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

.PHONY: all build run tidy clean
