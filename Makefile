hello:
	echo "Hello"
# Define the run-server target
run-server:
	go run ./cmd/main.go

tidy:
	go mod tidy
.PHONY: all
all: run-server