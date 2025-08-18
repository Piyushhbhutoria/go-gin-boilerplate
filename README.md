# Go Gin Boilerplate

> A production-ready starter project with Golang, Gin framework, and PostgreSQL

A modern, scalable Go web application boilerplate built with the Gin framework and PostgreSQL. Features clean architecture, database migrations, structured logging, authentication middleware, and environment-based configuration.

![](header.jpg)

## Features

- 🚀 **Gin Framework** - Fast HTTP web framework for Go
- 🗄️ **PostgreSQL** - Robust relational database with migrations
- 🔐 **Authentication Middleware** - JWT-based auth system
- 📝 **Structured Logging** - Configurable logging with different levels
- 🐳 **Docker Support** - Containerized development environment
- 🔧 **Environment Config** - Multiple configuration environments (dev, stage, prod)
- 📊 **Health Checks** - Built-in health monitoring endpoints
- 🏗️ **Clean Architecture** - Well-organized project structure
- 🚦 **Graceful Shutdown** - Proper server lifecycle management

## Project Structure

```
.
├── Makefile                 # Build and development commands
├── README.md               # This file
├── docker-compose.yml      # Docker development environment
├── go.mod                  # Go module dependencies
├── go.sum                  # Go module checksums
├── main.go                 # Application entry point
├── controllers/            # HTTP request handlers
│   └── health.go          # Health check endpoint
├── data/                   # Static data files
│   └── dummy.json         # Sample data
├── db/                     # Database configuration
│   └── migrations/         # Database migration files
│       ├── 000001_init.up.sql
│       ├── 000001_init.down.sql
│       └── 000001_init.sql
├── logger/                 # Logging configuration
│   └── logger.go          # Logger setup and utilities
├── middlewares/            # HTTP middleware
│   └── auth.go            # Authentication middleware
├── schema/                 # Database schema definitions
│   └── db.go              # Database connection and setup
├── server/                 # HTTP server configuration
│   ├── router.go          # Route definitions
│   └── server.go          # Server setup and configuration
├── store/                  # Data access layer
│   └── store.go           # Repository interfaces and implementations
└── util/                   # Utility functions
    └── http.go            # HTTP helper functions
```

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

## API Endpoints

### Health Check

```sh
curl http://localhost:3000/health
```

### Authentication (if implemented)

```sh
# Login
curl -X POST http://localhost:3000/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username":"user","password":"pass"}'

# Protected endpoint
curl http://localhost:3000/protected \
  -H "Authorization: Bearer <your-jwt-token>"
```

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
