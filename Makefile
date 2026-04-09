include makefiles/user.mk
include makefiles/docker.mk
.PHONY: fmt

fmt:
	go fmt all
