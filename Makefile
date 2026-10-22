# Common project tasks. Adjust per project; keep CLAUDE.md "Common commands" in sync.

.PHONY: help up down logs backend frontend test migrate migrate-down migrate-status fmt lint check hooks-install

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-12s\033[0m %s\n", $$1, $$2}'

up: ## Start the full stack (docker-compose)
	docker compose up --build

down: ## Stop the stack
	docker compose down

logs: ## Tail stack logs
	docker compose logs -f

backend: ## Run backend locally
	cd backend && go run ./cmd/api

frontend: ## Run frontend dev server
	cd frontend && yarn dev

test: ## Run all tests
	cd backend && go test ./...
	cd frontend && yarn test

migrate: ## Apply pending DB migrations
	cd backend && go run ./cmd/migrate up

migrate-down: ## Roll back the last migration group
	cd backend && go run ./cmd/migrate down

migrate-status: ## Show applied / pending migrations
	cd backend && go run ./cmd/migrate status

fmt: ## Format code
	cd backend && go fmt ./...
	cd frontend && yarn format

lint: ## Lint code
	cd backend && go vet ./...
	cd frontend && yarn lint

check: ## Run the AI Project OS health check (memory freshness, decay)
	bash scripts/check-system.sh

hooks-install: ## Install git hooks (pre-commit, commit-msg) from .githooks
	git config core.hooksPath .githooks
	@echo "Git hooks installed (core.hooksPath=.githooks)."
