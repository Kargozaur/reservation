.PHONY: docker-ubuild up down

docker-ubuild:
	docker build -t user-service services/user-service

up:
	docker-compose up -d

down:
	docker-compose down
