SHELL=/bin/sh

BASE_PATH       = ./
MAIN_GO_PATH    = /main.go

default: run

install:
	go mod download

run: install run_http_server

run_http_server:
	go run main.go http_server_run

test:
	go test ./...

test_force:
	go clean -testcache && go test ./...

test_force_v:
	go clean -testcache && go test -v .//...

test_coverage:
	go test -cover .//...

fmt:
	go fmt ./...

tidy:
	go mod tidy

lint:
	golangci-lint run --timeout 60m

swag_init:
	swag init --parseDependency --parseInternal -d $(BASE_PATH) -g $(MAIN_GO_PATH)

migrate_up:
	go run main.go migrate_up

migrate_down:
	go run main.go migrate_down

sync_data:
	go run main.go sync_data
