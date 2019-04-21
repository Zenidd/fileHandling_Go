# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build

# Binary names
BINARY_NAME=mybinary

build:
		$(GOBUILD) -o bin/$(BINARY_NAME) src/hello.go 