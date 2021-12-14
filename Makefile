# Should be set to the project name as it declared in go.mod!
APP=konkero-project/mark1

# Optionally set these args as environment vars in the shell. You
# could also pass them as parameters of `make`.
# For example: make build CMD=console
CMD?=git-repos
FLAGS?=-v
CLEANUP?=

default: lint test

# Optional includes that depend on project:
-include .env
-include doc.mk

# export env vars from .env file if exists
export $(shell if [ -f .env ]; then sed 's/=.*//' .env ; fi )

run:
	env
	./git-repos

build:
	go build $(FLAGS) $(APP)/cmd/$(CMD)

build-all:
	$(foreach dir,$(wildcard cmd/*), go build $(FLAGS) $(APP)/$(dir);)

build-race:
	go build $(FLAGS) -race $(APP)/cmd/$(CMD)

docker:
	docker build -f Dockerfile -t cryptoboyio/$(APP) .

lint:
	golangci-lint run -v ./...

test:
	go test $(FLAGS) ./...

test-race:
	go test $(FLAGS) -race ./...

gen:
	go generate $(FLAGS) ./...

mod:
	go mod tidy

clean:
	@echo $(CLEANUP)
	$(foreach f,$(CLEANUP),rm -rf $(f);)

.PHONY: build build-race build-all test test-race lint clean gen mod
