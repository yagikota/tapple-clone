include .env

.PHONY: help
help: ## 使い方
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: run-db
run-db: ## DB作成
	docker-compose up --build -d mysql

.PHONY: run-go-amd64
run-go-amd64: ## 起動
	docker-compose up --build server-amd64

# M1 Mac User
.PHONY: run-go-arm64
run-go-arm64: ## 起動
	docker-compose up --build server-arm64

.PHONY: down
down: ## down
	docker-compose down
