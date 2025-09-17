package main

import (
	"net/http"
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
)

type User struct {
	gorm.Model
	GitHubID       int64  `json:"github_id"`
	Username       string `json:"username"`
	Email          string `json:"email"`
	ProfilePicture string `json:"profile_picture"`
	AccessToken    string `json:"-"`
}

func GetCurrentUser(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authenticated"})
		return
	}

	var user User

	result := DB.First(&user, userID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":             user.ID,
        "github_id":      user.GitHubID,
        "username":       user.Username,
        "email":          user.Email,
        "profile_picture": user.ProfilePicture,
	})
}

func GetUserStats(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var stats struct {
		TotalTodos int64 `json:"total_todos"`
		CompletedTodos int64 `json:"completed_todos"`
		PendingTodos int64 `json:"pending_todos"`
	}

	if err := DB.Model(&Todo{}).Where("user_id = ?", userID).Count(&stats.TotalTodos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch stats for total todos"})
		return
	}

	if err := DB.Model(&Todo{}).Where("user_id = ? AND completed = ?", userID, true).Count(&stats.CompletedTodos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch stats for completed todos"})
		return
	}

	if err := DB.Model(&Todo{}).Where("user_id = ? AND completed = ?", userID, false).Count(&stats.PendingTodos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch stats for pending todos"})
		return
	}

	c.JSON(http.StatusOK, stats)
}