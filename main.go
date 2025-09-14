package main

import (
	"log"
	"github.com/gin-gonic/gin"
)

func main() {

	DB = ConnectDB()

	router := gin.Default()

	router.GET("/todos", GetTodos)
	router.POST("/todos", CreateTodos)
	router.GET("/todos/:id", GetTodoByID)
	router.PATCH("/todos/:id", UpdateTodo)
	router.DELETE("/todos/:id", DeleteTodo)

	log.Println("Server starting on http://localhost:8080")
	router.Run(":8080")
}