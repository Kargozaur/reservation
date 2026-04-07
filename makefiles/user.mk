.PHONY: ubuser

ubuser:
	go build -o usersvc services/user-service/main.go
