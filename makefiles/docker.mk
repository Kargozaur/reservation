.PHONY: docker-ubuild up down create_db

docker-ubuild:
	docker build -t user-service services/user_service

up:
	docker-compose up -d

down:
	docker-compose down

create_db:
	docker exec -it postgres bash ./docker/init.sh
