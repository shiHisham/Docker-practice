# Make watch docker containers status
watch:
	while true; do clear; docker ps; sleep 3; done

# Development environment
up-dev:
	docker-compose -f docker-compose.yaml -f docker-compose.dev.yaml up --build -d

up-golang-only-dev:
	docker-compose -f docker-compose.yaml -f docker-compose.dev.yaml up --build -d postgres api-golang

up-node-only-dev:
	docker-compose -f docker-compose.yaml -f docker-compose.dev.yaml up --build -d postgres api-node

up-react-only-dev:
	docker-compose -f docker-compose.yaml -f docker-compose.dev.yaml up --build -d postgres api-react

# Production
up-prod:
	docker-compose -f docker-compose.yaml -f docker-compose.prod.yaml up --build -d

up-golang-only-prod:
	docker-compose -f docker-compose.yaml -f docker-compose.prod.yaml up --build -d postgres api-golang

up-node-only-prod:
	docker-compose -f docker-compose.yaml -f docker-compose.prod.yaml up --build -d postgres api-node

up-react-only-prod:
	docker-compose -f docker-compose.yaml -f docker-compose.prod.yaml up --build -d postgres api-react

# Clean everything
clean:
	docker-compose -f docker-compose.yaml -f docker-compose.dev.yaml down -v
	docker-compose -f docker-compose.yaml -f docker-compose.prod.yaml down -v