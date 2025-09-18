package main

import (
	_ "go-todo-api/docs"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Todo API
// @version 1.0
// @description A Todo API with GitHub OAuth authentication and JWT tokens
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	log.Printf("GITHUB_CLIENT_ID: '%s'", os.Getenv("GITHUB_CLIENT_ID"))
	log.Printf("GITHUB_REDIRECT_URL: '%s'", os.Getenv("GITHUB_REDIRECT_URL"))

	InitOAuthConfig()

	jwtSecret = []byte(os.Getenv("JWT_SECRET_KEY"))
	if len(jwtSecret) == 0 {
		log.Fatal("JWT_SECRET_KEY is not set in .env file!")
	}

	DB = ConnectDB()
	DB.AutoMigrate(&Todo{}, &User{})

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:5173", "http://localhost:8080"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
