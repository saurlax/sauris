.PHONY: fmt check-fmt test vet build check

fmt:
	gofmt -w cmd internal

check-fmt:
	@test -z "$$(gofmt -l cmd internal)" || (gofmt -l cmd internal; exit 1)

test:
	go test ./...

vet:
	go vet ./...

build:
	go build -o bin/sauris ./cmd/sauris
	go build -o bin/sauris-server ./cmd/sauris-server
	go build -o bin/sauris-agent ./cmd/sauris-agent

check: check-fmt test vet build
