.PHONY: build clean install test

# Binary name
BINARY_NAME=norgannon

# Build the project
build:
	go build -o $(BINARY_NAME) .

# Clean the binary
clean:
	go clean
	rm -f $(BINARY_NAME)

# Install the binary to $GOPATH/bin
install: build
	go install

# Run tests
test:
	go test -v ./...

# Build for all major platforms
build-all:
	GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME)-linux-amd64 .
	GOOS=windows GOARCH=amd64 go build -o $(BINARY_NAME)-windows-amd64.exe .
	GOOS=darwin GOARCH=amd64 go build -o $(BINARY_NAME)-darwin-amd64 .

# Default target
all: build 