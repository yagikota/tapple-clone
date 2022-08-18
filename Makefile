include .env

GO_BIN:=$(shell go env GOPATH)/bin
WD:=$(shell pwd)
MYSQL_INFO:=-h 127.0.0.1 -P 3306 -u root
DB_NAME:=tapple_c
DML_DIR:=./migration/dml

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
	mysql $(MYSQL_INFO) $(DB_NAME) < $(DML_DIR)/dummy_users.up.sql
	mysql $(MYSQL_INFO) $(DB_NAME) < $(DML_DIR)/dummy_rooms.up.sql
	mysql $(MYSQL_INFO) $(DB_NAME) < $(DML_DIR)/dummy_room_users.up.sql
	mysql $(MYSQL_INFO) $(DB_NAME) < $(DML_DIR)/dummy_messages.up.sql
	mysql $(MYSQL_INFO) $(DB_NAME) < $(DML_DIR)/dummy_user_profile_images.up.sql

# ローカルデータDLETEのコマンド
.PHONY: delete
delete: ## delete
	mysql $(MYSQL_INFO) $(DB_NAME) < $(DML_DIR)/dummy_messages.down.sql
	mysql $(MYSQL_INFO) $(DB_NAME) < $(DML_DIR)/dummy_room_users.down.sql
	mysql $(MYSQL_INFO) $(DB_NAME) < $(DML_DIR)/dummy_rooms.down.sql
	mysql $(MYSQL_INFO) $(DB_NAME) < $(DML_DIR)/dummy_users.down.sql
	mysql $(MYSQL_INFO) $(DB_NAME) < $(DML_DIR)/dummy_user_profile_images.down.sql

# sqlboilerでmodel自動生成
.PHONY: generate-model
generate-model:
	@$(GO_BIN)/sqlboiler mysql

# コードを整形
.PHONY: fmt
fmt:
	gofmt -l -w .

# mock作成
# usecase
# serivice
# repository
.PHONY: generate-mock
generate-mock:
	@$(GO_BIN)/mockgen -source pkg/usecase/user_usecase.go -destination pkg/mock/usecase/user_usecase.go -package mock
	@$(GO_BIN)/mockgen -source pkg/domain/service/user_service.go -destination pkg/mock/service/user_service.go -package mock
	@$(GO_BIN)/mockgen -source pkg/domain/repository/user_repository.go -destination pkg/mock/repository/user_repository.go -package mock
