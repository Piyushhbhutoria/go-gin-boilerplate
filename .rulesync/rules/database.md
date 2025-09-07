---
root: true
targets: ["*"]
description: "Database usage rules and best practices for GORM integration"
globs: ["**/*.go", "**/migrations/*.sql"]
---

# Database Usage Rules

## GORM Integration Guidelines

### Database Access

- **Always use the store layer**: Access database through `store.GetDB()` instead of direct GORM connections
- **Never create new GORM instances**: Use the singleton pattern established in `store/store.go`
- **Import schema models**: Use models from `schema/db.go` package for all database operations

### Model Definitions

- **Location**: All database models MUST be defined in `schema/db.go`
- **GORM tags**: Use proper GORM tags for validation, indexing, and relationships
- **Soft deletes**: Use `gorm.DeletedAt` for models that need soft deletion capability
- **Timestamps**: Include `CreatedAt` and `UpdatedAt` fields for audit trails
- **JSON tags**: Always include JSON tags for API serialization

### Database Operations

- **Error handling**: Always check and handle GORM errors properly
- **Transactions**: Use GORM transactions for operations that must succeed or fail together
- **Preloading**: Use `Preload()` for loading related models to avoid N+1 queries
- **Selective loading**: Use `Select()` to load only required fields for performance

### Migration Strategy

- **GORM Auto-migrations**: **DISABLED** - Use golang-migrate instead
- **SQL Migrations**: **REQUIRED** for all schema changes using golang-migrate
- **Migration files**: Place SQL migrations in `db/migrations/` directory
- **Naming convention**: Use `make new name=descriptive_name` for new migrations
- **Migration commands**: Use `make up` to apply migrations, `make down` to rollback

### Schema Change Workflow

1. **Update GORM models** in `schema/db.go` first
2. **Create migration files** using `make new name=description`
3. **Write SQL** in the `.up.sql` file to match GORM model changes
4. **Write rollback SQL** in the `.down.sql` file
5. **Apply migration** using `make up`
6. **Update database.dbml** to reflect schema changes
7. **Test the application** to ensure everything works

### Query Best Practices

- **Parameterized queries**: Always use parameterized queries to prevent SQL injection
- **Indexing**: Add appropriate indexes using GORM tags (`gorm:"index"`)
- **Foreign keys**: Define relationships using `gorm:"foreignKey:FieldName"`
- **Constraints**: Use GORM tags for NOT NULL, UNIQUE, and other constraints

### Performance Guidelines

- **Batch operations**: Use `CreateInBatches()` for bulk inserts
- **Query optimization**: Use `Explain()` to analyze query performance
- **Connection pooling**: Rely on GORM's built-in connection pooling
- **Lazy loading**: Be mindful of when related models are loaded

### Security Rules

- **Environment variables**: Never hardcode database credentials
- **DATABASE_URL**: Always use the `DATABASE_URL` environment variable
- **Input validation**: Validate all inputs before database operations
- **SQL injection**: GORM handles this, but avoid raw SQL when possible

### Error Handling

- **Logging**: Use the project's logger for database errors
- **Graceful degradation**: Handle database connection failures gracefully
- **Retry logic**: Implement retry logic for transient database errors
- **Context usage**: Use context for timeout and cancellation

### Testing

- **Test database**: Use separate test database for unit tests
- **Migrations**: Run `make up` in test setup to apply all migrations
- **Cleanup**: Clean up test data after each test
- **Mocking**: Mock database operations for unit tests when appropriate
- **Migration testing**: Test both up and down migrations work correctly

## Code Examples

### Correct Usage

```go
// Get database instance
db := store.GetDB()

// Create with proper error handling
user := schema.User{
    Email: "user@example.com",
    Name:  "John Doe",
}
if err := db.Create(&user).Error; err != nil {
    logger.LogMessage("error", "failed to create user: %v", err)
    return err
}

// Query with preloading
var users []schema.User
if err := db.Preload("Orders").Find(&users).Error; err != nil {
    logger.LogMessage("error", "failed to fetch users: %v", err)
    return err
}
```

### Migration Example

```bash
# Create new migration
make new name=add_user_phone_field

# This creates:
# db/migrations/20250906173002_add_user_phone_field.up.sql
# db/migrations/20250906173002_add_user_phone_field.down.sql
```

```sql
-- up.sql
ALTER TABLE users ADD COLUMN phone VARCHAR(20);

-- down.sql  
ALTER TABLE users DROP COLUMN phone;
```

```bash
# Apply migration
make up

# Rollback if needed
make down
```

### Incorrect Usage

```go
// DON'T: Create new GORM instance
db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})

// DON'T: Ignore errors
db.Create(&user)

// DON'T: Use raw SQL without necessity
db.Raw("SELECT * FROM users WHERE email = ?", email).Scan(&user)

// DON'T: Rely on GORM auto-migration
// Auto-migration is disabled - use golang-migrate instead
```
