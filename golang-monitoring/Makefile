.PHONY: help build up down logs clean test

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

test: ## Test the application endpoints
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

setup: ## Initial setup - create directories and files
	mkdir -p logs
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
	@if command -v open >/dev/null 2>&1; then \
		open http://localhost:3000 & \
		open http://localhost:9090 & \
		open http://localhost:9093 & \
	fi
