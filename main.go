package main

import (
	"log"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	log.Printf("GITHUB_CLIENT_ID: '%s'", os.Getenv("GITHUB_CLIENT_ID"))
	log.Printf("GITHUB_CLIENT_SECRET: '%s'", os.Getenv("GITHUB_CLIENT_SECRET"))
	log.Printf("GITHUB_REDIRECT_URL: '%s'", os.Getenv("GITHUB_REDIRECT_URL"))

	InitOAuthConfig()

	DB = ConnectDB()
	DB.AutoMigrate(&Todo{}, &User{})

	router := gin.Default()

	router.GET("/auth/github", GithubLoginHandler)
	router.GET("/auth/github/callback", GithubCallbackHandler)

	router.GET("/todos", GetTodos)
	router.POST("/todos", CreateTodos)
	router.GET("/todos/:id", GetTodoByID)
	router.PATCH("/todos/:id", UpdateTodo)
	router.DELETE("/todos/:id", DeleteTodo)

	log.Println("Server starting on http://localhost:8080")
	router.Run(":8080")
}