PROG := chirpy
MIGRATE ?= status

go: run

all: clean goose sqlc run

run:
	@clear
	@echo "Runing ${PROG}..."
	@go run .

build:
	@echo "Builing ${PROG}..."
	@go build

clean:
	@echo "Cleaning ${PROG}..."
	@rm -f ${PROG}

goose:
	@echo "goose ${MIGRATE}..."
	@goose postgres "postgres://postgres:postgres@localhost:5432/chirpy?sslmode=disable" --dir ./sql/schema ${MIGRATE} 

sqlc:
	@echo "sqlc generate..."
	@sqlc generate