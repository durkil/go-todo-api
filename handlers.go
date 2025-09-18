package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var DB *gorm.DB

// GetTodos godoc
// @Summary Get all todos
// @Description Get all todos for the authenticated user
// @Tags todos
// @Security BearerAuth
// @Produce json
// @Success 200 {array} TodoResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/todos [get]
func GetTodos(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var todos []Todo
	if err := DB.Where("user_id = ?", userID).Find(&todos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch todos"})
		return
	}

	c.JSON(http.StatusOK, todos)
}

// CreateTodos godoc
// @Summary Create a new todo
// @Description Create a new todo for the authenticated user
// @Tags todos
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param todo body TodoRequest true "Todo data"
// @Success 201 {object} TodoResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/todos [post]
func CreateTodos(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var newTodo Todo
	if err := c.BindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if strings.TrimSpace(newTodo.Title) == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title can not be empty"})
		return
	}

	if len(strings.TrimSpace(newTodo.Title)) > 255 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title is too long"})
		return
	}

	newTodo.UserID = userID.(uint)
	newTodo.Title = strings.TrimSpace(newTodo.Title)

	if err := DB.Create(&newTodo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create todo"})
		return
	}
	c.JSON(http.StatusCreated, newTodo)
}

// GetTodoByID godoc
// @Summary Get todo by ID
// @Description Get a specific todo by ID for authenticated user
// @Tags todos
// @Security BearerAuth
// @Produce json
// @Param id path int true "Todo ID"
// @Success 200 {object} TodoResponse
// @Failure 404 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Router /api/todos/{id} [get]
func GetTodoByID(c *gin.Context) {
	userID, _ := c.Get("userID")
	id := c.Param("id")

	var todo Todo

	if err := DB.Where("user_id = ? AND id = ?", userID, id).First(&todo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	c.JSON(http.StatusOK, todo)
}

// UpdateTodo godoc
// @Summary Update todo
// @Description Update an existing todo for authenticated user
// @Tags todos
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "Todo ID"
// @Param todo body TodoRequest true "Updated todo data"
// @Success 200 {object} TodoResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Router /api/todos/{id} [patch]
func UpdateTodo(c *gin.Context) {
	userID, _ := c.Get("userID")
	id := c.Param("id")

	var updatedTodo Todo

	if err := c.BindJSON(&updatedTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var todo Todo

	if err := DB.Where("user_id = ? AND id = ?", userID, id).First(&todo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Todo not found"})
		return
	}

	DB.Model(&todo).Updates(&updatedTodo)
	c.JSON(http.StatusOK, todo)
}

// DeleteTodo godoc
// @Summary Delete todo
// @Description Delete a todo by ID for authenticated user
// @Tags todos
// @Security BearerAuth
// @Param id path int true "Todo ID"
// @Success 204
// @Failure 404 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Router /api/todos/{id} [delete]
func DeleteTodo(c *gin.Context) {
	userID, _ := c.Get("userID")
	id := c.Param("id")

	result := DB.Where("user_id = ? AND id = ?", userID, id).Delete(&Todo{}, id)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	c.Status(http.StatusNoContent)
}
