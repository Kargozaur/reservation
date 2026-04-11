.PHONY: buser

buser:
	go build -o usersvc services/user_service/main.go
