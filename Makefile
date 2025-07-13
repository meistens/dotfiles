.PHONY: help build up down logs clean test setup dev monitor migrate-create migrate-up migrate-down migrate-force migrate-version

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

build: ## Build the Go application
	go mod tidy
	go build -o bin/app cmd/main.go

up: ## Start all services
	podman-compose up -d
	@echo "Services starting up..."
	@echo "Grafana: http://localhost:3000 (admin/admin)"
	@echo "Prometheus: http://localhost:9090"
	@echo "Go App: http://localhost:8080"
	@echo "AlertManager: http://localhost:9093"

down: ## Stop all services
	podman-compose down

restart: down up ## Restart all services

logs: ## Show logs for all services
	podman-compose logs -f

logs-app: ## Show logs for Go application only
	podman-compose logs -f go-app

logs-prometheus: ## Show logs for Prometheus only
	podman-compose logs -f prometheus

logs-grafana: ## Show logs for Grafana only
	podman-compose logs -f grafana

logs-loki: ## Show logs for Loki only
	podman-compose logs -f loki

logs-promtail: ## Show logs for Promtail only
	podman-compose logs -f promtail

status: ## Show status of all services
	podman-compose ps

clean: ## Clean up containers and volumes
	podman-compose down -v
	podman system prune -f

test: ## Quickly test the application endpoints
	@echo "Testing health endpoint..."
	curl -f http://localhost:8080/health || echo "Health check failed"
	@echo ""
	@echo "Testing user creation..."
	curl -X POST -H "Content-Type: application/json" \
		-d '{"name":"Test User","email":"test@example.com"}' \
		http://localhost:8080/users || echo "User creation failed"
	@echo ""
	@echo "Testing metrics endpoint..."
	curl -f http://localhost:8080/metrics | head -20 || echo "Metrics endpoint failed"

setup: ## Initial setup - create directories and files. TODO: start here
	mkdir -p logs migrations
	mkdir -p grafana/provisioning/datasources
	mkdir -p grafana/dashboards
	@echo "apiVersion: 1\ndatasources:\n  - name: Prometheus\n    type: prometheus\n    access: proxy\n    url: http://prometheus:9090\n    isDefault: true\n  - name: Loki\n    type: loki\n    access: proxy\n    url: http://loki:3100" > grafana/provisioning/datasources/datasources.yml
	@echo "Setup complete!"

dev: ## Development mode - rebuild and restart
	podman-compose down
	podman-compose build --no-cache go-app
	podman-compose up -d

monitor: ## Open monitoring dashboards
	@echo "Opening monitoring dashboards..."
	@echo "Grafana: http://localhost:3000"
	@echo "Prometheus: http://localhost:9090"
	@echo "AlertManager: http://localhost:9093"

# Migration targets - Using postgres superuser for migrations
migrate-create: ## Create a new migration file (usage: make migrate-create NAME=create_users_table)
	@if [ -z "$(NAME)" ]; then \
		echo "Error: NAME is required. Usage: make migrate-create NAME=create_users_table"; \
		exit 1; \
	fi
	@mkdir -p migrations
	@migrate create -ext sql -dir migrations -seq $(NAME)
	@echo "Migration files created in migrations/ directory"

migrate-up: ## Run all up migrations.TODO: change, review others
	@migrate -path migrations -database "postgres://postgres:postgres@localhost:5432/monitoring_testing?sslmode=require" up

migrate-down: ## Run one down migration
	@migrate -path migrations -database "postgres://postgres:postgres@localhost:5432/monitoring_testing?sslmode=require" down 1

migrate-force: ## Force migration version (usage: make migrate-force VERSION=1)
	@if [ -z "$(VERSION)" ]; then \
		echo "Error: VERSION is required. Usage: make migrate-force VERSION=1"; \
		exit 1; \
	fi
	@migrate -path migrations -database "postgres://postgres:postgres@localhost:5432/monitoring_testing?sslmode=require" force $(VERSION)

migrate-version: ## Show current migration version
	@migrate -path migrations -database "postgres://postgres:postgres@localhost:5432/monitoring_testing?sslmode=require" version
