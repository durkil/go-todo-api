package main

import (
	"net/http"
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
)

var DB *gorm.DB

func GetTodos(c *gin.Context) {
	userID, _ := c.Get("userID")

	var todos []Todo
	DB.Where("user_id = ?", userID).Find(&todos)
	c.JSON(http.StatusOK, todos)
}

func CreateTodos(c *gin.Context) {
	userID, _ := c.Get("userID")

	var newTodo Todo
	if err := c.BindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newTodo.UserID = userID.(uint)
	DB.Create(&newTodo)
	c.JSON(http.StatusCreated, newTodo)
}

func GetTodoByID(c *gin.Context) {
	userID, _ := c.Get("userID")
	id := c.Param("id")

	var todo Todo

	result := DB.Where("user_id = ? AND id = ?", userID, id).First(&todo)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	c.JSON(http.StatusOK, todo)
}

func UpdateTodo(c *gin.Context) {
	userID, _ := c.Get("userID")
	id := c.Param("id")

	var updatedTodo Todo

	if err := c.BindJSON(&updatedTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var todo Todo
	result := DB.Where("user_id = ? AND id = ?", userID, id).First(&todo)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Todo not found"})
		return
	}

	DB.Model(&todo).Updates(&updatedTodo)
	c.JSON(http.StatusOK, todo)
}

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
