
include .env

create_network:
	docker network create mesh-istio-network

remove_network:
	docker network rm -f mesh-istio-network

up:
	docker-compose --env-file .env -f docker-compose.yml -p $(PROJECT_NAME) up -d

down:
	docker-compose --env-file .env -f docker-compose.yml -p $(PROJECT_NAME) down

build:
	cd src && docker build --target prod -t linhhnbkdn/auth-miroservice:$(VERSION) . --no-cache

build_and_push: build
	docker push linhhnbkdn/auth-miroservice:$(VERSION)
