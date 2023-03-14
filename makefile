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
