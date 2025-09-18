package main

import "time"

// TodoResponse represents a todo item in API responses
// @Description Todo item information for API responses
type TodoResponse struct {
	ID        uint      `json:"id" example:"1"`
	Title     string    `json:"title" example:"Complete project"`
	Completed bool      `json:"completed" example:"false"`
	UserID    uint      `json:"user_id" example:"1"`
	CreatedAt time.Time `json:"created_at" example:"2023-01-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2023-01-01T00:00:00Z"`
}

// TodoRequest represents a request to create or update a todo
// @Description Request body for creating or updating a todo
type TodoRequest struct {
	Title     string `json:"title" binding:"required" example:"Complete project"`
	Completed *bool  `json:"completed,omitempty" example:"false"`
}

// UserResponse represents a user in API responses
// @Description User account information for API responses
type UserResponse struct {
	ID             uint      `json:"id" example:"1"`
	GitHubID       int64     `json:"github_id" example:"12345678"`
	Username       string    `json:"username" example:"johndoe"`
	Email          string    `json:"email" example:"john@example.com"`
	ProfilePicture string    `json:"profile_picture" example:"https://github.com/johndoe.png"`
	CreatedAt      time.Time `json:"created_at" example:"2023-01-01T00:00:00Z"`
	UpdatedAt      time.Time `json:"updated_at" example:"2023-01-01T00:00:00Z"`
}

// ErrorResponse represents an error response
// @Description Error response
type ErrorResponse struct {
	Error string `json:"error" example:"Invalid input"`
}

// AuthResponse represents authentication response
// @Description Authentication response with JWT token
type AuthResponse struct {
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
}
