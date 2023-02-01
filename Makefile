GOCMD=go
GOBUILD=$(GOCMD) build

all: build
build:
	$(GOBUILD) -o ./bin/broker ./cmd/broker
	$(GOBUILD) -o ./bin/consumer ./cmd/consumer
	$(GOBUILD) -o ./bin/producer ./cmd/producer

clean:
	rm -rf ./bin

