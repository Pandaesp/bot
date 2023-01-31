include .env
export

.PHONY: run
run: 
	go run cmd/bot/main.go