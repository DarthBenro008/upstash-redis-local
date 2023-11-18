PROJECT_NAME = "upstash-redis-local"
BASE=$(shell pwd)
BUILD_DIR=$(BASE)/bin
VERSION ?= v1.0
BUILD_DATE = $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
COMMIT_SHA = $(shell git rev-parse --short HEAD)
LDFLAGS = -ldflags="-X main.Version=${VERSION}"
PACKAGE = $(shell go list -m)

.PHONY: clean
clean:
	@echo "> Cleaning Build targets"
	rm -rf bin

.PHONY: deps
deps:
	@echo "> Installing dependencies"
	@go mod tidy
	@go mod download

.PHONY: build
build: deps
	@echo "> Building upstash-redis-local backend Server Binary"
	go build ${LDFLAGS} -o ${BUILD_DIR}/${PROJECT_NAME}
	@echo "> Binary has been built successfully"


.PHONY: run
run: build
	@echo "> Running ${PROJECT_NAME}"
	${BUILD_DIR}/${PROJECT_NAME}