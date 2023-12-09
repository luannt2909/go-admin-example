BRACKET_OPEN := (
BRACKET_CLOSE := )

MESSAGE_CONV := $(subst -, ,${TAG})
PROJECT_NAME := $(patsubst bump$(BRACKET_OPEN)%$(BRACKET_CLOSE):,%,$(word 2,${MESSAGE_CONV}))
SERVICE := $(word 3,${MESSAGE_CONV})
VERSION := $(lastword ${MESSAGE_CONV})
PROJECT := $(PROJECT_NAME)/$(SERVICE)
LDFLAGS += -X go-reminder-bot.BuildVersion=$(shell git describe --tags --always | sed 's/-/+/' | sed 's/^v//') -X go-reminder-bot.BuildCommit=$(shell git rev-parse HEAD) -X go-reminder-bot.BuildDate=$(shell date +'%Y-%m-%dT%H:%M:%S%Z')

BUILD_CHECK_FILE_NAME := "buildcheck"
BUILD_CHECK_FILES_LIST = $(shell find . -name 'main.go' | grep -v /vendor/ | sort)
PKG_LIST := $(shell go list ./... | grep -v /vendor/)
LINT_DIR_LIST = $(shell ls -d */ | grep -v -E scripts\|webs\|vendor\|tools\|cache/)
# The following line is for running golangci-lint on specific dir
#LINT_DIR_LIST = $(shell ls -d */ | grep <dirname>/)

.PHONY: all build clean test coverage coverhtml lint

all: buildcheck

buildcheck: ## Try to build all main.go
	$(eval INDEX = 0)
	$(eval COUNT = $(words $(BUILD_CHECK_FILES_LIST)))
	@echo ╔═══ START BUILD PROCESS
	@for file in $(BUILD_CHECK_FILES_LIST); do \
		INDEX=$$(($${INDEX}+1)); \
		echo ╠═ Building [$$INDEX/$(COUNT)] $$(dirname $$file); \
		go build -o ${BUILD_CHECK_FILE_NAME} $$(dirname $$file); \
	done
	@rm ${BUILD_CHECK_FILE_NAME}
	@echo ╚═══ FINISH BUILD PROCESS

lint: ## Golangci-lint
	$(eval INDEX = 0)
	$(eval COUNT = $(words $(LINT_DIR_LIST)))
	@echo ╔═══ START LINTING
	@for dir in $(LINT_DIR_LIST); do \
		INDEX=$$(($${INDEX}+1)); \
		echo ╠═ Checking [$$INDEX/$(COUNT)] $$dir; \
		golangci-lint run ./$$dir/...; \
	done
	@echo ╚═══ FINISH LINTING

lintsingle: ## Golangci-lint
	golangci-lint run -v ./${pkg}/...

test: ## Run unittests
	@go test -race ${PKG_LIST}

race: ## Run data race detector
	@go test -race ${PKG_LIST}

#msan: dep ## Run memory sanitizer
# @go test -msan -short ${UNITTEST_PKG_LIST}

coverage: ## Generate global code coverage report
	./scripts/coverage.sh;

coverhtml:
	go-acc -o coverage.out ./${pkg}/...
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out -o coverage.html

test-coverage:
	go test -coverprofile=coverage.out -covermode count ./...
	go tool cover -func=coverage.out

docker-build: ## build docker
	docker build -t go-reminder-bot -f Dockerfile .
	docker tag go-reminder-bot luannt2909/go-reminder-bot:latest
	docker push luannt2909/go-reminder-bot:latest

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
