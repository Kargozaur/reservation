UDIR=services/user_service/migrations
USERDSN=host=localhost user=$(POSTGRES_USER) port=5433 password=$(POSTGRES_PASSWORD) dbname=$(UDBNAME) sslmode=disable

.PHONY: migrate-usersvc-create migrate-usersvc-up migrate-usersvc-down


migrate-usersvc-create:
	@read -p "Enter migration name: " name; \
	goose -dir $(UDIR) create $$name sql

migrate-usersvc-up:
	goose -dir $(UDIR) "$(USERDSN)" up

migrate-usersvc-down:
	goose -dir $(UDIR) "$(USERDSN)" down
