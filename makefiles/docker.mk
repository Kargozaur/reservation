.PHONY: docker-ubuild

docker-ubuild:
	docker build -t user-service services/user-service
