GO_CMD = go
BINARY_NAME = ${service}-main.out

all: build test

test:
	$(GO_CMD) test -v ./...

build:
	$(GO_CMD) build -o ${BINARY_NAME} ./$(service)/cmd/main.go

