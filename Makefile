BINARY_NAME=server

BIN_DIR=bin

SRC_DIR=cmd/api/main.go

TEST_DIR=./...

.PHONY: all build clean run test deps

all: build

build:
	go build -o $(BIN_DIR)/$(BINARY_NAME) $(SRC_DIR)

clean:
	rm -rf $(BIN_DIR)

run: build
	./$(BIN_DIR)/$(BINARY_NAME)

test:
	go test $(TEST_DIR)

deps:
	go mod tidy

update-deps:
	go get -u ./...
	go mod tidy

migration:
	@migrate create -ext sql -dir cmd/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down

migrate-force:
	@migrate -path cmd/migrations -database "mysql://root:root@tcp(localhost:3306)/golang_db" force $(version)