# Database and migrations setup

## Requirements
- Docker + Docker Compose
- Homebrew

## Install migrate CLI
brew install golang-migrate

## Start Postgres
# Creates a local Postgres 16 instance on port 5432
# Data persists in a Docker volume

docker compose up -d postgres

## Configure environment
cp .env.example .env
# If you already use direnv, .envrc sets a default DATABASE_URL too

## Run migrations
# Create new migration files
make new name=create_users_table

# Apply all up migrations
make up

# Show current version
make status

# Rollback last migration
make down

