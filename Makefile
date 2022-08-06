include .env

GO_BIN:=$(shell go env GOPATH)/bin

.PHONY: help
help: ## 使い方
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: run-db
run-db: ## DB作成
	docker-compose up --build -d mysql

.PHONY: run-go
run-go: ## 起動
	docker-compose up --build server

.PHONY: down
down: ## down
	docker-compose down

# TODO: golang-migrateのmigrationコマンド作成

# sqlboilerでmodel自動生成
.PHONY: generate-model
generate-model:
	@$(GO_BIN)/sqlboiler mysql

# コードを整形
.PHONY: fmt
fmt:
	gofmt -l -w .