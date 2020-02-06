build:
	go build -o ./build/namegrind ./cmd/main.go

test:
	go test ./... --race

.PHONY: build test