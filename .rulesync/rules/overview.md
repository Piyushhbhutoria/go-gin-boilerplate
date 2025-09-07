---
root: true
targets: ["*"]
description: "Go Gin Boilerplate project overview and development guidelines"
globs: ["**/*"]
---

# Go Gin Boilerplate Project Overview

This is a RESTful API boilerplate built with Go, Gin framework, GORM, and PostgreSQL.

## Tech Stack

- **Backend**: Go with Gin web framework
- **Database**: PostgreSQL with GORM ORM
- **Migrations**: golang-migrate (not GORM auto-migrations)
- **Documentation**: Swagger/OpenAPI with swag
- **Environment**: direnv for environment management

## Project Structure

```
├── controllers/     # HTTP handlers and business logic
├── schema/         # Database models and response types
├── store/          # Database connection and configuration
├── server/         # Router and server setup
├── middlewares/    # Custom middleware (auth, CORS, etc.)
├── logger/         # Logging utilities
├── util/           # Utility functions
├── db/migrations/  # Database migration files
└── docs/           # Auto-generated Swagger documentation
```

## General Guidelines

- Use Go for all backend code
- Follow Go naming conventions (camelCase for private, PascalCase for public)
- Write self-documenting code with clear variable and function names
- Use meaningful comments for complex business logic
- Always handle errors explicitly
- Use proper HTTP status codes

## Code Style

- Use tabs for indentation (Go standard)
- Use meaningful variable and function names
- Group imports: standard library, third-party, local packages
- Use `gofmt` for code formatting
- Follow Go idioms and best practices

## Architecture Principles

- **Controllers**: Handle HTTP requests and responses
- **Schema**: Define data models and API response types
- **Store**: Manage database connections and operations
- **Middleware**: Handle cross-cutting concerns (auth, logging, CORS)
- **Separation of Concerns**: Keep business logic separate from HTTP handling
- **Error Handling**: Always check and handle errors properly

## Database Guidelines

- Use GORM for database operations
- Use golang-migrate for schema changes (NOT GORM auto-migrations)
- Define models in `schema/db.go`
- Create migration files using `make new name=description`
- Apply migrations with `make up`

## API Documentation

- Add Swagger annotations to all controller methods
- Use realistic examples in response models
- Document all parameters, responses, and error cases
- Regenerate docs with `make swagger` after changes
- Access documentation at `http://localhost:3000/swagger/index.html`

## Development Workflow

1. **Start database**: `docker compose up -d postgres`
2. **Set environment**: Configure `DATABASE_URL` in `.env`
3. **Run migrations**: `make up`
4. **Build application**: `make build`
5. **Generate docs**: `make swagger`
6. **Never run server during development** - only build for syntax checking
