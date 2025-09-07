# Go Gin Boilerplate

> A production-ready starter project with Golang, Gin framework, and PostgreSQL

A modern, scalable Go web application boilerplate built with the Gin framework and PostgreSQL. Features clean architecture, database migrations, structured logging, authentication middleware, and environment-based configuration.

![](header.jpg)

## Features

- 🚀 **Gin Framework** - Fast HTTP web framework for Go
- 🗄️ **PostgreSQL** - Robust relational database with GORM ORM
- 🔄 **Database Migrations** - golang-migrate for schema management
- 📚 **Swagger Documentation** - Interactive API documentation with realistic examples
- 🔐 **Authentication Middleware** - JWT-based auth system
- 📝 **Structured Logging** - Configurable logging with different levels
- 🐳 **Docker Support** - Containerized development environment
- 🔧 **Environment Config** - Multiple configuration environments (dev, stage, prod)
- 📊 **Health Checks** - Built-in health monitoring endpoints
- 🏗️ **Clean Architecture** - Well-organized project structure
- 🚦 **Graceful Shutdown** - Proper server lifecycle management
- 🤖 **AI Agent Rules** - .rulesync configuration for consistent AI assistance
- 📋 **Realistic API Responses** - Production-ready response models with examples
- 🔍 **Search & Pagination** - Built-in filtering and pagination for endpoints

## Project Structure

```text
├── .env.example
├── .envrc
├── .github/
│   └── workflows/
│       ├── build.yml
│       ├── go.yaml
│       └── lint.yml
├── .gitignore
├── .travis.yml
├── Makefile
├── README.md
├── README_DB.md
├── controllers/
│   └── health.go
├── data/
│   └── dummy.json
├── db/
│   └── migrations/
│       ├── 000001_init.down.sql
│       ├── 000001_init.sql
│       └── 000001_init.up.sql
├── docker-compose.yml
├── go.mod
├── go.sum
├── header.jpg
├── logger/
│   └── logger.go
├── main.go
├── middlewares/
│   └── auth.go
├── run
├── schema/
│   └── db.go
├── server/
│   ├── router.go
│   └── server.go
├── store/
│   └── store.go
└── util/
    └── http.go
├── .rulesync/
│   └── rules/
│       └── overview.md
└── docs/
    ├── README.md
    ├── docs.go
    ├── swagger.json
    └── swagger.yaml
```

## .rulesync Configuration

This project uses `.rulesync` to maintain consistent coding standards across AI coding agents. The `.rulesync/rules/` directory contains project-specific guidelines.

### Quick Setup

```sh
# Install and initialize
npm install -g @rulesync/cli

# Sync with your AI agents
npx rulesync import --targets cursor
npx rulesync generate --targets cursor
```

### Documentation

For detailed setup and usage instructions, see the [.rulesync documentation](https://rulesync.dev/docs).

## Prerequisites

- Go 1.25+
- PostgreSQL 16+
- Docker & Docker Compose (optional)

## Installation

1. **Clone the repository**

   ```sh
   git clone <your-repo-url>
   cd go-gin-boilerplate
   ```

2. **Set up environment variables**

   ```sh
   # Copy the existing .envrc file (already configured for local development)
   cp .envrc .env
   
   # Or create your own .env file with:
   echo "DATABASE_URL=postgres://postgres:postgres@localhost:5432/postgres" > .env
   ```

3. **Start PostgreSQL (using Docker)**

   ```sh
   docker-compose up -d postgres
   ```

4. **Run database migrations**

   ```sh
   make up
   ```

## Development

### Using Makefile (Recommended)

```sh
# Database operations
make up          # Run all up migrations
make down        # Roll back the last migration step
make status      # Print current migration version
make new name=create_users_table  # Create new migration

# Help
make help        # Show all available targets
```

### Running the Application

```sh
# Option 1: Use the run script (includes tests and build)
./run

# Option 2: Manual build and run
go build -tags=jsoniter -o build/application
./build/application

# Option 3: Direct run
go run main.go
```

### Manual Commands

```sh
# Install dependencies
go mod download

# Run tests
go test ./...

# Build with JSON iter tags
go build -tags=jsoniter -o build/application

# Run database migrations manually
migrate -database "$DATABASE_URL" -path db/migrations up
```

### Using Docker

```sh
# Start PostgreSQL only
docker-compose up -d postgres

# Start the application (after building)
./run
```

## Configuration

The application uses environment variables for configuration:

- `.envrc` - Default environment configuration (already included)
- `.env` - Your local environment overrides

Key configuration options:

- `DATABASE_URL` - PostgreSQL connection string
- `PORT` - Server port (default: 3000)
- `SERVICE_NAME` - Application service name

### Environment Variables

```sh
# Database
DATABASE_URL=postgres://postgres:postgres@localhost:5432/postgres

# Server
PORT=3000
SERVICE_NAME=go-gin-boilerplate
```

## API Documentation

This project includes Swagger/OpenAPI documentation with interactive UI.

### Quick Access

```sh
# Generate documentation
make swagger

# Access Swagger UI
http://localhost:3000/swagger/index.html
```

### Available Endpoints

- `GET /health` - Health check
- `GET /users` - Get users (with pagination & search)
- `POST /users` - Create user
- `GET /users/{id}` - Get user by ID

### Documentation

For detailed API documentation and examples, see [docs/README.md](docs/README.md).

## Database

### Migrations

Database migrations are managed using the `migrate` tool through Makefile targets:

```sh
# Apply migrations
make up

# Rollback migrations
make down

# Check migration status
make status

# Create new migration
make new name=migration_name

# Force set migration version
make force v=1
```

### Schema

Database schema is defined in `schema/db.go` and includes:

- Connection pooling
- Migration support
- Transaction handling

## Logging

Structured logging is configured in `logger/logger.go` with support for:

- Multiple log levels (DEBUG, INFO, WARN, ERROR)
- JSON formatting
- File and console output
- Request ID tracking

## Testing

```sh
# Run tests (included in ./run script)
go test ./...

# Run tests with verbose output
go test ./controllers -v

# Run tests before building (./run does this automatically)
go test ./...
go build -tags=jsoniter -o build/application
```

Built with ❤️ using Go and Gin
