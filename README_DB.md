# Database Setup and GORM Integration

This Go Gin boilerplate includes both traditional SQL migrations and GORM (Go Object-Relational Mapping) for database operations.

## Requirements

- Docker
- Homebrew

## Database Setup

### Install migrate CLI

```bash
brew install golang-migrate
```

### Start Postgres

Creates a local Postgres 16 instance on port 5432. Data persists in a Docker volume.

```bash
docker compose up -d postgres
```

### Configure environment

```bash
cp .env.example .env
# If you already use direnv, .envrc sets a default DATABASE_URL too
```

## Migration Management

### Traditional SQL Migrations

For complex schema changes or when you need precise control over SQL:

```bash
# Create new migration files
make new name=create_users_table

# Apply all up migrations
make up

# Show current version
make status

# Rollback last migration
make down
```

### GORM Integration

GORM is used for database operations but migrations are handled by golang-migrate. The application will not auto-migrate on startup.

### Dependencies

- `gorm.io/gorm` - Core GORM library
- `gorm.io/driver/postgres` - PostgreSQL driver for GORM

### Database Models

Located in `schema/db.go`:

- **User** - User management with email, name, and timestamps
- **Product** - Product catalog with name, description, price, and stock
- **Order** - Order management with user relationship
- **OrderItem** - Order line items with product relationships

### Updated Store Layer

- Replaced raw `database/sql` with GORM
- Automatic database migrations on startup
- Improved connection handling and logging

### Example API Endpoints

- `GET /users` - List all users
- `POST /users` - Create a new user
- `GET /users/:id` - Get user by ID

## Usage

### Environment Variables

```bash
export DATABASE_URL="postgres://username:password@localhost/dbname?sslmode=disable"
```

### Running the Application

```bash
go run main.go
```

### Database Operations with GORM

```go
// Get GORM instance
db := store.GetDB()

// Create a user
user := schema.User{
    Email: "user@example.com",
    Name:  "John Doe",
}
db.Create(&user)

// Find users
var users []schema.User
db.Find(&users)

// Find with conditions
var user schema.User
db.Where("email = ?", "user@example.com").First(&user)
```

## Model Features

- **Soft Deletes** - Models use `gorm.DeletedAt` for soft deletion
- **Timestamps** - Automatic `CreatedAt` and `UpdatedAt` fields
- **Relationships** - Foreign key relationships between models
- **Validation** - GORM tags for field validation and constraints

## Migration Strategy

### When to use SQL Migrations

- Complex schema changes requiring custom SQL
- Data migrations or transformations
- Performance-critical index creation
- Database-specific optimizations

### When to use GORM Auto-Migrations

- **Not recommended** - Use golang-migrate instead
- GORM auto-migrations are disabled in this project
- All schema changes should go through proper migration files

## Development Workflow

1. **Start the database**: `docker compose up -d postgres`
2. **Set environment**: Configure `DATABASE_URL` in `.env`
3. **Run migrations**: `make up` (run database migrations)
4. **Run application**: `go run main.go`
5. **For schema changes**: Create new migrations with `make new name=description`
