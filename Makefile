build:
	@go build -o bin/wc cmd/main.go

run: build
	@./bin/wc

test:
	@go test -v ./... -short
