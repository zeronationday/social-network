include .env  
export

run:
	@go run ./cmd

up:
	@goose up

down:
	@goose down
