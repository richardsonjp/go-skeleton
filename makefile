api-server:
	go run ./cmd/apiserver/main.go server

api-config:
	go run ./cmd/apiserver/main.go config

build-api-server:
	go build ./cmd/apiserver/main.go

sample-get-todo-list:
	go run ./sample/main.go gtl

sample-create-todo-list:
	go run ./sample/main.go cti

migrate-up:
	include ./.env
	export
	migrate -path ./pkg/clients/db/migration -database "postgresql://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}${DB_PORT}?sslmode=disable" -verbose up

migrate-down:
	include ./.env
	export
	migrate -path ./pkg/clients/db/migration -database "postgresql://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}${DB_PORT}?sslmode=disable" -verbose down
