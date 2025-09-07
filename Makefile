SHELL := /bin/zsh

# Load .env if present
ifneq (,$(wildcard ./.env))
  include .env
  export $$(shell sed -n 's/\s*#.*$$//; s/\s*export\s\+//; s/\s*=.*$$//p' .env)
endif

MIGRATIONS_DIR := db/migrations
DB_URL ?= $(DATABASE_URL)

.PHONY: help up down force version new createdb dropdb psql status swagger swagger-clean build run

help:
	@echo "Available targets:"
	@echo "  up        - run all up migrations"
	@echo "  down      - roll back the last migration step"
	@echo "  force     - set database version without running migrations (v=)"
	@echo "  version   - print current migration version"
	@echo "  new       - create a new timestamped migration name=your_title"
	@echo "  status    - print migration status"
	@echo "  swagger   - generate Swagger documentation"
	@echo "  swagger-clean - clean generated Swagger files"
	@echo "  build     - build the application"
	@echo "  run       - run the application"

up:
	@if [ -z "$(DB_URL)" ]; then echo "DATABASE_URL not set"; exit 1; fi
	migrate -database "$(DB_URL)" -path $(MIGRATIONS_DIR) up

status:
	@if [ -z "$(DB_URL)" ]; then echo "DATABASE_URL not set"; exit 1; fi
	migrate -database "$(DB_URL)" -path $(MIGRATIONS_DIR) version || true

version: status

down:
	@if [ -z "$(DB_URL)" ]; then echo "DATABASE_URL not set"; exit 1; fi
	migrate -database "$(DB_URL)" -path $(MIGRATIONS_DIR) down 1

force:
	@if [ -z "$(v)" ]; then echo "Usage: make force v=<version>"; exit 1; fi
	@if [ -z "$(DB_URL)" ]; then echo "DATABASE_URL not set"; exit 1; fi
	migrate -database "$(DB_URL)" -path $(MIGRATIONS_DIR) force $(v)

new:
	@if [ -z "$(name)" ]; then echo "Usage: make new name=create_users_table"; exit 1; fi
	@ts=$$(date +%Y%m%d%H%M%S); \
	up="$(MIGRATIONS_DIR)/$${ts}_$(name).up.sql"; \
	down="$(MIGRATIONS_DIR)/$${ts}_$(name).down.sql"; \
	mkdir -p $(MIGRATIONS_DIR); \
	touch "$$up" "$$down"; \
	echo "Created $$up and $$down"

swagger:
	@swag init -g main.go -o docs
	@echo "Swagger documentation generated in docs/"
	@echo "Access documentation at: http://localhost:<port>/swagger/index.html"

swagger-clean:
	@rm -f docs/docs.go docs/swagger.json docs/swagger.yaml
	@echo "Swagger files cleaned"

build:
	go build -o build/application .
	echo "Application built successfully"

run:
	echo "Starting application..."
	./run

