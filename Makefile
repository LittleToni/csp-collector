init:
	cp .env.example .env

build:
	go build -o csp-collector ./cmd/csp-collector

start:
	./csp-collector

test:
	go test ./...
