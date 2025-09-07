package controllers

import (
	"net/http"
	"strconv"

	"github.com/Piyushhbhutoria/go-gin-boilerplate/schema"
	"github.com/Piyushhbhutoria/go-gin-boilerplate/store"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct{}

// GetUsers godoc
// @Summary Get all users
// @Description Retrieve a paginated list of all users with optional filtering
// @Tags users
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of users per page" default(10)
// @Param search query string false "Search by name or email"
// @Success 200 {object} schema.UsersListResponse "List of users"
// @Failure 400 {object} schema.ErrorResponse "Invalid pagination parameters"
// @Failure 500 {object} schema.ErrorResponse "Internal server error"
// @Router /users [get]
func (uc UserController) GetUsers(c *gin.Context) {
	// Parse pagination parameters
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	search := c.Query("search")

	// Validate pagination parameters
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	offset := (page - 1) * limit

	var users []schema.User
	var totalCount int64
	db := store.GetDB()

	// Build query with optional search
	query := db.Model(&schema.User{})
	if search != "" {
		query = query.Where("name ILIKE ? OR email ILIKE ?", "%"+search+"%", "%"+search+"%")
	}

	// Get total count
	if err := query.Count(&totalCount).Error; err != nil {
		c.JSON(http.StatusInternalServerError, schema.ErrorResponse{
			Error:   "Failed to count users",
			Code:    "USER_COUNT_ERROR",
			Details: err.Error(),
		})
		return
	}

	// Get users with pagination
	if err := query.Offset(offset).Limit(limit).Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, schema.ErrorResponse{
			Error:   "Failed to fetch users",
			Code:    "USER_FETCH_ERROR",
			Details: err.Error(),
		})
		return
	}

	// Convert to response format
	userResponses := make([]schema.UserResponse, len(users))
	for i, user := range users {
		userResponses[i] = schema.UserResponse{
			ID:        user.ID,
			Email:     user.Email,
			Name:      user.Name,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}
	}

	response := schema.UsersListResponse{
		Users: userResponses,
		Count: int(totalCount),
		Page:  page,
		Limit: limit,
	}

	c.JSON(http.StatusOK, response)
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the provided information
// @Tags users
// @Accept json
// @Produce json
// @Param user body schema.CreateUserRequest true "User information"
// @Success 201 {object} schema.CreateUserResponse "User created successfully"
// @Failure 400 {object} schema.ValidationErrorResponse "Validation failed"
// @Failure 409 {object} schema.ErrorResponse "Email already exists"
// @Failure 500 {object} schema.ErrorResponse "Internal server error"
// @Router /users [post]
func (uc UserController) CreateUser(c *gin.Context) {
	var req schema.CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, schema.ErrorResponse{
			Error:   "Invalid request format",
			Code:    "INVALID_JSON",
			Details: err.Error(),
		})
		return
	}

	// Check if email already exists
	db := store.GetDB()
	var existingUser schema.User
	if err := db.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, schema.ErrorResponse{
			Error:   "Email already exists",
			Code:    "EMAIL_EXISTS",
			Details: "A user with this email address already exists",
		})
		return
	}

	// Create new user
	user := schema.User{
		Email: req.Email,
		Name:  req.Name,
	}

	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, schema.ErrorResponse{
			Error:   "Failed to create user",
			Code:    "USER_CREATE_ERROR",
			Details: err.Error(),
		})
		return
	}

	// Convert to response format
	userResponse := schema.UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	response := schema.CreateUserResponse{
		User: userResponse,
		// Token would be generated here in a real authentication system
		// Token: generateJWTToken(user.ID),
	}

	c.JSON(http.StatusCreated, response)
}

// GetUser godoc
// @Summary Get user by ID
// @Description Retrieve a specific user by their ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID" example(1)
// @Success 200 {object} schema.UserResponse "User information"
// @Failure 400 {object} schema.ErrorResponse "Invalid user ID"
// @Failure 404 {object} schema.ErrorResponse "User not found"
// @Failure 500 {object} schema.ErrorResponse "Internal server error"
// @Router /users/{id} [get]
func (uc UserController) GetUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, schema.ErrorResponse{
			Error:   "Invalid user ID",
			Code:    "INVALID_USER_ID",
			Details: "User ID must be a valid number",
		})
		return
	}

	var user schema.User
	db := store.GetDB()

	if err := db.First(&user, uint(id)).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, schema.ErrorResponse{
				Error:   "User not found",
				Code:    "USER_NOT_FOUND",
				Details: "No user exists with the provided ID",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, schema.ErrorResponse{
			Error:   "Failed to fetch user",
			Code:    "USER_FETCH_ERROR",
			Details: err.Error(),
		})
		return
	}

	// Convert to response format
	userResponse := schema.UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	c.JSON(http.StatusOK, userResponse)
}
