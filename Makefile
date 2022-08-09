include .env

GO_BIN:=$(shell go env GOPATH)/bin
WD:=$(shell pwd)
DB_NAME:=tapple_c

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

# テーブル作成
.PHONY: migrate
migrate: ## migrate
	migrate -path migration/ddl/ -database 'mysql://root:@tcp(localhost:3306)/$(DB_NAME)?parseTime=true&loc=Local' up

# ローカルデータ挿入のコマンド
.PHONY: seed
seed: ## seed
	mysql -h 127.0.0.1 -P 3306 -u root tapple_c < migration/dml/dummy_users.up.sql
	mysql -h 127.0.0.1 -P 3306 -u root tapple_c < migration/dml/dummy_rooms.up.sql
	mysql -h 127.0.0.1 -P 3306 -u root tapple_c < migration/dml/dummy_room_users.up.sql
	mysql -h 127.0.0.1 -P 3306 -u root tapple_c < migration/dml/dummy_messages.up.sql

# ローカルデータDLETEのコマンド
.PHONY: delete
delete: ## delete
	mysql -h 127.0.0.1 -P 3306 -u root tapple_c < migration/dml/dummy_messages.down.sql
	mysql -h 127.0.0.1 -P 3306 -u root tapple_c < migration/dml/dummy_room_users.down.sql
	mysql -h 127.0.0.1 -P 3306 -u root tapple_c < migration/dml/dummy_rooms.down.sql
	mysql -h 127.0.0.1 -P 3306 -u root tapple_c < migration/dml/dummy_users.down.sql

# sqlboilerでmodel自動生成
.PHONY: generate-model
generate-model:
	@$(GO_BIN)/sqlboiler mysql

# コードを整形
.PHONY: fmt
fmt:
	gofmt -l -w .
