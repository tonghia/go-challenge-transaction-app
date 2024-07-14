SHELL=/bin/bash -o pipefail

GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)
GOFMT := "goimports"


# Define the source folder and the mock destination folder
SOURCE_PB_FOLDER := ./pb
MOCK_PB_DESTINATION := ./pb

# List all the source files with the "_gprc.pb.go" suffix
SOURCE_PB_FILES := $$(find $(SOURCE_PB_FOLDER) -maxdepth 1 -name '*_grpc.pb.go' -printf '%f\n')

# Define the mockgen command
MOCKGEN_CMD := "mockgen"

HOST = 127.0.0.1
PORT = 3306
DATABASE = user_transaction
USER = root
PASSWORD = secret

# Build 
VCS_REF := $(shell git rev-parse HEAD)
BUILD_DATE := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

.PHONY: install
install:
	go install github.com/bufbuild/buf/cmd/buf@latest
	go install go.uber.org/mock/mockgen@latest
	go install golang.org/x/tools/cmd/goimports@latest
	go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

.PHONY: fmt
fmt: ## Run gofmt for all .go files
	@$(GOFMT) -w $(GOFMT_FILES)

.PHONY: generate
generate: ## Generate buf
	@rm -f pb/*
	@rm -f docs/*
	@rm -f **/mock/*
	@make gogen
	@buf generate proto
	@make genpbmocks
	@$(GOFMT) -w pb/*

.PHONY: gogen
gogen: ## Run go generate for whole project
	@go generate ./...

genpbmocks: ## Generate mock for grpc client
	@for filename in $(SOURCE_PB_FILES); do \
		mockfilename=$(MOCK_PB_DESTINATION)/$${filename/.pb.go/.pb.mock.go} ; \
		mockgen --source="$(SOURCE_PB_FOLDER)/$$filename" -destination="$$mockfilename" -package=pb ; \
		echo "Generated mock for $$filename as $$mockfilename" ; \
	done

cleangen: ## Clean generated files
	@rm -f pb/*
	@rm -f docs/*
	@rm -f **/mock/*

cleanmock: ## Clean generated mock files
	@rm -f **/mock/*

unit-test: ## Run unit test and coverage
	go test -v -timeout 5m -coverprofile dist/cover.out ./...
	go tool cover -html=dist/cover.out -o dist/cover.html

test: ## Run go test for whole project
	@go test -v ./...

lint: ## Run linter
	@golangci-lint run ./...

.PHONY: migrate-up
migrate-up: ## Run migrations
	migrate -database "mysql://$(USER):$(PASSWORD)@tcp($(HOST):$(PORT))/$(DATABASE)" -path migrations up

.PHONY: migrate-down
migrate-down: ## Rollback migrations
	migrate -database "mysql://$(USER):$(PASSWORD)@tcp($(HOST):$(PORT))/$(DATABASE)" -path migrations down

.PHONY: compose-up
compose-up: ## Start server
	DOCKER_BUILDKIT=1 docker build -f Dockerfile --build-arg=COMMIT=$(VCS_REF) --build-arg=BUILD_DATE=$(BUILD_DATE) -t user_transaction_app .
	docker compose -f compose.yaml up -d

.PHONY: compose-down
compose-down: ## Stop server
	docker compose -f compose.yaml down

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
