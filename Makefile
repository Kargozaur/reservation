include .env
include db.env
export

include makefiles/user.mk
include makefiles/docker.mk
include makefiles/migrations.mk

.PHONY: fmt

fmt:
	go fmt all
