include .env

# set default shell
SHELL = bash -e -o pipefail

default: run

.PHONY:	install
install:
	go mod tidy
	go mod vendor

.PHONY:	lint
lint:
	golangci-lint run 

.PHONY:	build
build:
	mkdir -p bin
	go build -o bin/gomora-dapp \
	    cmd/main.go

.PHONY:	build-dev
build-dev:
	mkdir -p bin
	go build -race -o bin/gomora-dapp \
	    cmd/main.go

.PHONY:	test
test:
	go test -race -v -p 1 ./...
	
.PHONY:	run
run:	build
	./bin/gomora-dapp

.PHONY: run-dev
run-dev: build-dev
	./bin/gomora-dapp

.PHONY: up
up:
	docker compose down
	docker compose up -d --build

.PHONY:	schema
schema:
	mkdir -p infrastructures/database/mysql/migrations
	migrate create -ext sql -dir infrastructures/database/mysql/migrations -seq ${NAME}

.PHONY: migrate-up
migrate-up: 
	migrate -path infrastructures/database/mysql/migrations -database "mysql://${DB_USERNAME}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_DATABASE}" -verbose up ${STEPS}

.PHONY: migrate-down
migrate-down:
	migrate -path infrastructures/database/mysql/migrations -database "mysql://${DB_USERNAME}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_DATABASE}" -verbose down ${STEPS}

.PHONY: migrate-version
migrate-version:
	migrate -path infrastructures/database/mysql/migration -database "mysql://${DB_USERNAME}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_DATABASE}" version

.PHONY: migrate-force
migrate-force:
	migrate -path infrastructures/database/mysql/migration -database "mysql://${DB_USERNAME}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_DATABASE}" force ${STEPS}

.PHONY: contract-build
contract-build:
	rm -rf infrastructures/smartcontracts/greeter/Greeter.go
	abigen --abi=infrastructures/smartcontracts/build/Greeter.abi --pkg=greeter --out=infrastructures/smartcontracts/greeter/Greeter.go
	