.PHONY: test
test:
	go test -race ./...

build:
	go build ./...

fmt:
	go fmt ./...

.PHONY: lint
lint:
	golangci-lint run ./...
