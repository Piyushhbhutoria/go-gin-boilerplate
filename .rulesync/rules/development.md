---
root: true
targets: ["*"]
description: "General development rules and workflow"
globs: ["**/*"]
---

# Development Rules

## Pre-Development Analysis

### 1. Database Schema Reference

- **ALWAYS refer to `@database.dbml`** before implementing any requirement
- Verify database schema matches the current implementation in `schema/db.go`
- Ensure all required tables, relationships, and constraints are properly defined
- **Update `database.dbml`** after any schema changes

## Development Workflow

### 3. API Endpoint Development

- **Add Swagger annotations** to all new controller methods
- **Use proper HTTP status codes** and response structures
- **Document all parameters** (path, query, body)
- **Include example responses** for success and error cases
- **Regenerate Swagger docs** after any API changes

### 2. Error Handling Standards

- **Check for errors** after every operation that can fail
- Use proper error wrapping with context
- Log errors using the project's logger system
- Return meaningful error messages to clients
- Handle database errors gracefully
- Implement proper HTTP status codes

### 3. Code Quality Checks

- Follow Go best practices and idioms
- Use meaningful variable and function names
- Keep functions small and focused
- Add comments for complex logic
- Follow the project's existing code style

## Post-Development Validation

### 4. Build Verification

- **Always check for successful build** after implementation
- Use `make build` to verify successful compilation
- Run `go mod tidy` to clean up dependencies
- **DO NOT run the server** - only verify build compilation

### 5. Documentation Updates

- **Update README.md** if new features or setup steps are added
- Update `README_DB.md` for database-related changes
- Add API documentation for new endpoints
- Update environment variable documentation
- Document any new configuration options
- **ALWAYS regenerate Swagger documentation** after adding/modifying API endpoints

## Implementation Checklist

### Before Starting

- [ ] Review `@database.dbml` for schema requirements
- [ ] Plan the implementation approach

### During Development

- [ ] Implement error handling for all operations
- [ ] Follow Go best practices
- [ ] Use proper logging
- [ ] Handle edge cases
- [ ] Add Swagger annotations for new API endpoints

### After Implementation

- [ ] Verify successful build (`make build`)
- [ ] Update `database.dbml` if schema changes were made
- [ ] Regenerate Swagger documentation (`make swagger`)
- [ ] Update documentation if needed
- [ ] Commit with GPG signature

## Code Examples

### Proper Error Handling

```go
// Database operation with error handling
func CreateUser(user *schema.User) error {
    db := store.GetDB()
    if err := db.Create(user).Error; err != nil {
        logger.LogMessage("error", "failed to create user: %v", err)
        return fmt.Errorf("failed to create user: %w", err)
    }
    return nil
}
```

### Build Verification

```bash
# Always run these commands after development
go mod tidy
make build  # Verify successful compilation
make swagger  # Regenerate Swagger documentation
```

### Swagger Documentation

```bash
# After adding/modifying API endpoints
make swagger

# Access documentation at: http://localhost:3000/swagger/index.html
# (Only after manually starting the server for testing)
```

### Swagger Annotation Example

```go
// GetUser godoc
// @Summary Get user by ID
// @Description Retrieve a specific user by their ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} map[string]interface{} "User information"
// @Failure 404 {object} map[string]string "User not found"
// @Router /users/{id} [get]
func (uc UserController) GetUser(c *gin.Context) {
    // Implementation
}
```

## Documentation Standards

### README Updates

- Add new features to the features list
- Update installation/setup instructions
- Document new environment variables
- Add usage examples for new functionality
- Update API documentation

### Code Comments

- Document public functions and types
- Explain complex business logic
- Add TODO comments for future improvements
- Document any workarounds or temporary solutions
