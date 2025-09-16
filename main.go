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

	jwtSecret = []byte(os.Getenv("JWT_SECRET_KEY"))
	if len(jwtSecret) == 0 {
		log.Fatal("JWT_SECRET_KEY is not set in .env file!")
	}

	DB = ConnectDB()
	DB.AutoMigrate(&Todo{}, &User{})

	router := gin.Default()

	router.GET("/auth/github", GithubLoginHandler)
	router.GET("/auth/github/callback", GithubCallbackHandler)

	api := router.Group("/api")
	api.Use(AuthMiddleware())
	{
		api.GET("/me", GetCurrentUser)
		api.GET("/todos", GetTodos)
		api.POST("/todos", CreateTodos)
		api.GET("/todos/:id", GetTodoByID)
		api.PATCH("/todos/:id", UpdateTodo)
		api.DELETE("/todos/:id", DeleteTodo)
	}

	log.Println("Server starting on http://localhost:8080")
	router.Run(":8080")
}