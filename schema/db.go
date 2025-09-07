package schema

import (
	"time"

	"gorm.io/gorm"
)

// User represents a user in the system
type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Email     string         `gorm:"uniqueIndex;not null" json:"email"`
	Name      string         `gorm:"not null" json:"name"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

// Product represents a product in the system
type Product struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Name        string         `gorm:"not null" json:"name"`
	Description string         `json:"description"`
	Price       float64        `gorm:"not null" json:"price"`
	Stock       int            `gorm:"default:0" json:"stock"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

// Order represents an order in the system
type Order struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	UserID    uint           `gorm:"not null" json:"user_id"`
	User      User           `gorm:"foreignKey:UserID" json:"user"`
	Total     float64        `gorm:"not null" json:"total"`
	Status    string         `gorm:"default:'pending'" json:"status"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

// OrderItem represents an item in an order
type OrderItem struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	OrderID   uint      `gorm:"not null" json:"order_id"`
	Order     Order     `gorm:"foreignKey:OrderID" json:"order"`
	ProductID uint      `gorm:"not null" json:"product_id"`
	Product   Product   `gorm:"foreignKey:ProductID" json:"product"`
	Quantity  int       `gorm:"not null" json:"quantity"`
	Price     float64   `gorm:"not null" json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UserResponse represents the user data in API responses
type UserResponse struct {
	ID        uint      `json:"id" example:"1"`
	Email     string    `json:"email" example:"john.doe@example.com"`
	Name      string    `json:"name" example:"John Doe"`
	CreatedAt time.Time `json:"created_at" example:"2023-12-01T10:30:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2023-12-01T10:30:00Z"`
}

// UsersListResponse represents the response for getting all users
type UsersListResponse struct {
	Users []UserResponse `json:"users"`
	Count int            `json:"count" example:"25"`
	Page  int            `json:"page" example:"1"`
	Limit int            `json:"limit" example:"10"`
}

// CreateUserRequest represents the request body for creating a user
type CreateUserRequest struct {
	Email string `json:"email" binding:"required,email" example:"jane.smith@example.com"`
	Name  string `json:"name" binding:"required,min=2,max=100" example:"Jane Smith"`
}

// CreateUserResponse represents the response for creating a user
type CreateUserResponse struct {
	User  UserResponse `json:"user"`
	Token string       `json:"token,omitempty" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
}

// ErrorResponse represents error responses
type ErrorResponse struct {
	Error   string `json:"error" example:"Failed to fetch users"`
	Code    string `json:"code,omitempty" example:"USER_FETCH_ERROR"`
	Details string `json:"details,omitempty" example:"Database connection timeout"`
}

// ValidationErrorResponse represents validation error responses
type ValidationErrorResponse struct {
	Error   string            `json:"error" example:"Validation failed"`
	Details map[string]string `json:"details" example:"email:invalid email format"`
}
