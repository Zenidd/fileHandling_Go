# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build

# Binary names
BINARY_NAME=fileHandling_Go

build:
		@$(GOBUILD) -o bin/$(BINARY_NAME) src/hello.go

clean:
		@printf "Cleaning binaries \n"
		@rm bin/*
