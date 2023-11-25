init:
	cp .env.example .env

build:
	go build

start:
	./csp-collector

test:
	go test
