.PHONY: default build test run docs clean

APP_NAME=cutlink

default: run

run:
	@docker compose up -d
	
up:
	docker-compose up --build --force-recreate -d

down:
	docker-compose down

run-local:
	docker-compose -f only-redis-compose.yaml up -d;go run cmd/main.go

redis-kill:
	@docker-compose -f only-redis-compose.yaml down
