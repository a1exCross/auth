include .env

LOCAL_BIN:=$(CURDIR)/bin

install-lint:
	mkdir -p bin
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.53.3

lint:
	$(LOCAL_BIN)/golangci-lint run ./... --config .golangci.pipeline.yaml

install-deps:
	mkdir -p bin
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.14.0

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

generate:
	make generate-user-api

generate-user-api:
	mkdir -p pkg/user_v1
	protoc --proto_path api/user_v1 \
	--go_out=pkg/user_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/user_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/user_v1/user.proto

go-build:
	go build -o ./build/auth_server cmd/server/main.go

docker-build:
	docker buildx build --no-cache --platform linux/amd64 -t cr.selcloud.ru/alexzabolotskikh/auth_server .

up-local:
	docker-compose --project-directory ./ -f config/local/docker-compose.yml up -d

up-prod:
	docker-compose --project-directory ./ -f config/prod/docker-compose.yml up -d

run-app-local: up-local
	go run cmd/server/main.go --config-path=config/local/.env

local-migration-status:
	$(LOCAL_BIN)/goose -dir internal/${MIGRATION_DIR} postgres ${PG_DSN} status -v

local-migration-up:
	$(LOCAL_BIN)/goose -dir internal/${MIGRATION_DIR} postgres ${PG_DSN} up -v

local-migration-down:
	$(LOCAL_BIN)/goose -dir internal/${MIGRATION_DIR} postgres ${PG_DSN} down -v

#run-app-prod: up-prod
#	go run cmd/server/main.go --config-path=config/prod/.env
