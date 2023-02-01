# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test

all: build
build:
	$(GOBUILD) -o ./bin/broker_demo ./cmd

clean:
	rm -rf ./bin
