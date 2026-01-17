include .env  
export

run:
	@go run ./cmd

migration:
	@goose create $(name) sql -s

up:
	@goose up

down:
	@goose down
