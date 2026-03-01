GO_CMD = go

all: build test

test:
	$(GO_CMD) test -v ./...

build:
	$(G0_CMD) build -o main ./$(service)/cmd/main.go

