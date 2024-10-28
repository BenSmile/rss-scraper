include .env

.PHONY: build
build:
	@go build -o rssaggregator

.PHONY: run
run: build
	@./rssaggregator

.PHONY: dev
dev:
	@reflex --start-service -r '\.go$$' make run

.PHONY: migrate-up
migrate-up:
	@goose -dir ./sql/schema postgres ${DB_URL} up

.PHONY: migrate-down
migrate-down:
	@goose -dir ./sql/schema postgres ${DB_URL} down

.PHONY: sqlc-gen
sqlc-gen:
	@sqlc generate