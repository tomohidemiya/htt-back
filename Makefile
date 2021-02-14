NAME=hitotsu-backend-v1
VERSION=0.0.1

export GO111MODULES=on

## Install Dependencies
.PHONY: deps
deps:
	go get -v -t -d ./...

# 開発用の依存関係解決
## Setup Dev env
.PHONY: deps
dev-deps: deps
	GO111MODULES-off \
	go get github.com/golang/lint/golint

.PHONY: build
## build: Compile the packages.
build:
	@go build -o $(NAME)

.PHONY: run
## run: Build and Run in development mode.
run: build
	@./$(NAME) -e development

.PHONY: run-prod
## run-prod: Build and Run in production mode.
run-prod: build
	@./$(NAME) -e production

.PHONY: lint
## lint: Run lint by go vet and golint
lint: dev-deps
	go vet ./...
	golint -set_exit_status ./...

.PHONY: test
## test: Run tests with verbose mode
test:
	go test -v ./...

.PHONY: migrate
migrate:
    @sql-migrate up

.PHONY: help
all: help
# help: show this help message
help: Makefile
	@echo
	@echo " Choose a command to run in "$(APP_NAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo
