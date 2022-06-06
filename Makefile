PACKAGES := $(shell go list ./... | grep -v /vendor/)

.PHONY: run
run:
	go run cmd/server/main.go

.PHONY: build
build:
	go build -o bin/server cmd/server/main.go
 
.PHONY: docker-build
docker-build:
	@docker build -t tg_translate_bot --no-cache .

.PHONY: docker-run
docker-run: 
	@docker run --name tg_bot --rm -p 8080:80 tg_translate_bot

.PHONY: delete-image
delete-image:
	@docker rmi tg_translate_bot

.PHONY: fmt
fmt: ## run "go fmt" on all Go packages
	@go fmt $(PACKAGES)
