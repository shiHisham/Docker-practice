# Development environment
up-dev:
	docker-compose -f docker-compose.yaml -f docker-compose.dev.yaml up --build

up-golang-only-dev:
	docker-compose -f docker-compose.yaml -f docker-compose.dev.yaml up --build -d postgres api-golang

up-node-only-dev:
	docker-compose -f docker-compose.yaml -f docker-compose.dev.yaml up --build postgres api-node

# Production
up-prod:
	docker-compose -f docker-compose.yaml -f docker-compose.prod.yaml up --build -d

up-golang-only-prod:
	docker-compose -f docker-compose.yaml -f docker-compose.prod.yaml up --build -d postgres api-golang

up-node-only-prod:
	docker-compose -f docker-compose.yaml -f docker-compose.prod.yaml up --build postgres api-node

# Clean everything
clean:
	docker-compose -f docker-compose.yaml -f docker-compose.dev.yaml down -v
	docker-compose -f docker-compose.yaml -f docker-compose.prod.yaml down -v
