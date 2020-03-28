# Params definitions
GOCMD=go
GORUN=$(GOCMD) run
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
BINARY_OUTPUT_NAME=tictacgoe

# Commands
default: 
	$(GORUN) cmd/tictacgoe/main.go

run: 
	$(GORUN) cmd/tictacgoe/main.go

build: 
	$(GOBUILD) -o $(BINARY_OUTPUT_NAME) -v cmd/tictacgoe/main.go

test: 
	$(GOTEST) ./... -cover ./... -v
