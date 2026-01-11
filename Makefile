.PHONY: help up down restart logs ps clean build test infra services all stop

# Default target
.DEFAULT_GOAL := help

# Auto-detect docker compose command
DOCKER_COMPOSE := $(shell command -v docker-compose 2> /dev/null)
ifndef DOCKER_COMPOSE
    DOCKER_COMPOSE := $(shell docker compose version > /dev/null 2>&1 && echo "docker compose")
endif

# Colors for terminal output
BLUE := \033[0;34m
GREEN := \033[0;32m
YELLOW := \033[1;33m
NC := \033[0m # No Color

## help: Display this help message
help:
	@echo "$(GREEN)========================================$(NC)"
	@echo "$(GREEN)  Dianping Microservices - Make Commands$(NC)"
	@echo "$(GREEN)========================================$(NC)"
	@echo ""
	@echo "$(YELLOW)Available commands:$(NC)"
	@grep -E '^## ' Makefile | sed 's/## /  /' | column -t -s ':'
	@echo ""

## up: Start all services
up:
	@echo "$(BLUE)Starting all services...$(NC)"
	$(DOCKER_COMPOSE) up -d
	@echo "$(GREEN)✓ All services started$(NC)"
	@$(MAKE) ps

## down: Stop all services
down:
	@echo "$(YELLOW)Stopping all services...$(NC)"
	$(DOCKER_COMPOSE) down
	@echo "$(GREEN)✓ Services stopped$(NC)"

## stop: Stop services without removing containers
stop:
	@echo "$(YELLOW)Stopping services...$(NC)"
	$(DOCKER_COMPOSE) stop
	@echo "$(GREEN)✓ Services stopped$(NC)"

## restart: Restart all services
restart:
	@echo "$(BLUE)Restarting all services...$(NC)"
	$(DOCKER_COMPOSE) restart
	@echo "$(GREEN)✓ Services restarted$(NC)"

## infra: Start only infrastructure services (MySQL, Redis, etcd)
infra:
	@echo "$(BLUE)Starting infrastructure services...$(NC)"
	$(DOCKER_COMPOSE) up -d mysql redis etcd
	@echo "$(GREEN)✓ Infrastructure services started$(NC)"
	@$(MAKE) ps

## services: Start RPC services and API gateway
services:
	@echo "$(BLUE)Starting RPC services and API gateway...$(NC)"
	$(DOCKER_COMPOSE) up -d user-rpc shop-rpc voucher-rpc order-rpc blog-rpc api-gateway
	@echo "$(GREEN)✓ Application services started$(NC)"
	@$(MAKE) ps

## monitoring: Start monitoring services (Prometheus, Grafana, Jaeger)
monitoring:
	@echo "$(BLUE)Starting monitoring services...$(NC)"
	$(DOCKER_COMPOSE) up -d prometheus grafana jaeger
	@echo "$(GREEN)✓ Monitoring services started$(NC)"
	@echo "Grafana:    http://localhost:3000 (admin/admin)"
	@echo "Prometheus: http://localhost:9090"
	@echo "Jaeger:     http://localhost:16686"

## logs: View logs for all services
logs:
	$(DOCKER_COMPOSE) logs -f

## logs-api: View API gateway logs
logs-api:
	$(DOCKER_COMPOSE) logs -f api-gateway

## logs-user: View user RPC logs
logs-user:
	$(DOCKER_COMPOSE) logs -f user-rpc

## logs-shop: View shop RPC logs
logs-shop:
	$(DOCKER_COMPOSE) logs -f shop-rpc

## logs-voucher: View voucher RPC logs
logs-voucher:
	$(DOCKER_COMPOSE) logs -f voucher-rpc

## logs-order: View order RPC logs
logs-order:
	$(DOCKER_COMPOSE) logs -f order-rpc

## logs-blog: View blog RPC logs
logs-blog:
	$(DOCKER_COMPOSE) logs -f blog-rpc

## ps: Show service status
ps:
	@echo "$(BLUE)Service Status:$(NC)"
	@$(DOCKER_COMPOSE) ps

## build: Build or rebuild services
build:
	@echo "$(BLUE)Building services...$(NC)"
	$(DOCKER_COMPOSE) build --no-cache
	@echo "$(GREEN)✓ Build complete$(NC)"

## rebuild: Rebuild and restart specific service (usage: make rebuild SERVICE=user-rpc)
rebuild:
	@if [ -z "$(SERVICE)" ]; then \
		echo "$(YELLOW)Usage: make rebuild SERVICE=<service-name>$(NC)"; \
		echo "Example: make rebuild SERVICE=user-rpc"; \
		exit 1; \
	fi
	@echo "$(BLUE)Rebuilding $(SERVICE)...$(NC)"
	$(DOCKER_COMPOSE) build --no-cache $(SERVICE)
	$(DOCKER_COMPOSE) up -d $(SERVICE)
	@echo "$(GREEN)✓ $(SERVICE) rebuilt and restarted$(NC)"

## clean: Stop and remove all containers, networks, and volumes
clean:
	@echo "$(YELLOW)⚠️  This will remove all containers, networks, and volumes!$(NC)"
	@read -p "Are you sure? [y/N] " -n 1 -r; \
	echo; \
	if [[ $$REPLY =~ ^[Yy]$$ ]]; then \
		$(DOCKER_COMPOSE) down -v; \
		rm -rf data/; \
		echo "$(GREEN)✓ Cleanup complete$(NC)"; \
	else \
		echo "$(YELLOW)Cleanup cancelled$(NC)"; \
	fi

## prune: Remove unused Docker resources
prune:
	@echo "$(YELLOW)Cleaning up unused Docker resources...$(NC)"
	docker system prune -f
	@echo "$(GREEN)✓ Cleanup complete$(NC)"

## mysql: Connect to MySQL
mysql:
	docker exec -it dianping-mysql mysql -uroot -proot hmdp

## redis: Connect to Redis CLI
redis:
	docker exec -it dianping-redis redis-cli

## etcd: Check etcd services
etcd:
	@echo "$(BLUE)Registered services in etcd:$(NC)"
	docker exec dianping-etcd etcdctl get "" --prefix --keys-only

## test: Run integration tests (placeholder)
test:
	@echo "$(YELLOW)Running integration tests...$(NC)"
	@echo "$(YELLOW)Test suite not implemented yet$(NC)"

## health: Check health status of all services
health:
	@echo "$(BLUE)Checking service health...$(NC)"
	@docker inspect --format='{{.Name}}: {{.State.Health.Status}}' $$($(DOCKER_COMPOSE) ps -q) 2>/dev/null || echo "No health checks configured"

## urls: Display service URLs
urls:
	@echo "$(GREEN)========================================$(NC)"
	@echo "$(GREEN)Service URLs:$(NC)"
	@echo "$(GREEN)========================================$(NC)"
	@echo "API Gateway:    http://localhost:8081"
	@echo "User RPC:       localhost:8001"
	@echo "Shop RPC:       localhost:8002"
	@echo "Voucher RPC:    localhost:8003"
	@echo "Order RPC:      localhost:8004"
	@echo "Blog RPC:       localhost:8005"
	@echo ""
	@echo "MySQL:          localhost:3306 (root/root)"
	@echo "Redis:          localhost:6379"
	@echo "etcd:           localhost:2379"
	@echo ""
	@echo "Grafana:        http://localhost:3000 (admin/admin)"
	@echo "Prometheus:     http://localhost:9090"
	@echo "Jaeger UI:      http://localhost:16686"
	@echo "$(GREEN)========================================$(NC)"

## init: Initialize environment (create .env from .env.example)
init:
	@if [ ! -f .env ]; then \
		cp .env.example .env; \
		echo "$(GREEN)✓ .env file created from .env.example$(NC)"; \
		echo "$(YELLOW)Please review and update .env file if needed$(NC)"; \
	else \
		echo "$(YELLOW).env file already exists$(NC)"; \
	fi

## all: Complete setup (init + infra + wait + services + monitoring)
all: init
	@echo "$(GREEN)========================================$(NC)"
	@echo "$(GREEN)  Complete Deployment$(NC)"
	@echo "$(GREEN)========================================$(NC)"
	@$(MAKE) infra
	@echo "$(YELLOW)Waiting for infrastructure to be ready...$(NC)"
	@sleep 30
	@$(MAKE) services
	@echo "$(YELLOW)Waiting for services to start...$(NC)"
	@sleep 20
	@$(MAKE) monitoring
	@echo ""
	@$(MAKE) urls
	@echo ""
	@echo "$(GREEN)✓ All services deployed successfully!$(NC)"
