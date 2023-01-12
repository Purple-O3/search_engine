all: build

build:
	go build -ldflags "-s -w" -o scripts/search_engine  cmd/main.go

run:
	cd scripts && ./search_engine

clean:
	cd scripts && rm ./search_engine

pull-pika:
	docker pull pikadb/pika:latest

.PHONY: all build install run clean
