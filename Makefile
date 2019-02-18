
.PHONY: build integration-test docker-up docker-db-up docker-down clear 

build:
	./build.sh

integration-test: docker-db-up
	@go test -v ./...

docker-up:
	@docker-compose up -d

docker-db-up:
	@docker-compose up -d db

docker-down:
	@docker-compose down

clear: docker-down