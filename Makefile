BRANCH := $(shell git branch --show-current)
COMMIT := $(shell git log -1 --format='%h')
DATE := $(shell date +%Y%m%d%H%M%S)
IMAGE_TAG := ${BRANCH}-${DATE}-${COMMIT}

all: build 

build:
	go build -ldflags "-s -w" -o scripts/search_engine  cmd/main.go

run:
	cd scripts && ./search_engine

clean:
	cd scripts && rm ./search_engine

image-builder:
	docker build --file scripts/Dockerfile --tag search_engine:${IMAGE_TAG} .

pull-pika:
	docker pull pikadb/pika:latest

.PHONY: all build install run clean
