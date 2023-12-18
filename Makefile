all: fmt vet test bin

.PHONY: fmt vet test

fmt:
	go fmt ./...

vet:
	go vet ./...

test:
	go test ./...

bin: bin/exinc

bin/exinc: fmt vet test
	go build -o bin/exinc ./cmd/exinc
