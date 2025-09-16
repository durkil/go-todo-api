package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

var githubOAuthConfig *oauth2.Config

func InitOAuthConfig() {
	githubOAuthConfig = &oauth2.Config{
		ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
		ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("GITHUB_REDIRECT_URL"),
		Scopes:       []string{"user:email"},
		Endpoint:     github.Endpoint,
	}

	log.Printf("OAuth Config Initialized:")
	log.Printf("ClientID: '%s'", githubOAuthConfig.ClientID)
	log.Printf("RedirectURL: '%s'", githubOAuthConfig.RedirectURL)

	if githubOAuthConfig.ClientID == "" {
		log.Fatal("GITHUB_CLIENT_ID is not set!")
	}
}

func GithubLoginHandler(c *gin.Context) {
	log.Println("=== GitHub Login Handler Called ===")

	if githubOAuthConfig == nil {
		log.Println("ERROR: OAuth config not initialized")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "OAuth not configured"})
		return
	}

	url := githubOAuthConfig.AuthCodeURL("state", oauth2.AccessTypeOnline)
	log.Printf("Generated OAuth URL: %s", url)

	c.Redirect(http.StatusTemporaryRedirect, url)
}

func GithubCallbackHandler(c *gin.Context) {
	log.Println("=== GitHub Callback Handler Called ===")

	code := c.Query("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Authorization code not found"})
		return
	}

	log.Printf("Received authorization code: %s...", code[:10])

	token, err := githubOAuthConfig.Exchange(c, code)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to exchange token: " + err.Error()})
		return
	}

	client := githubOAuthConfig.Client(c, token)
	resp, err := client.Get("https://api.github.com/user")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get user info: " + err.Error()})
		return
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response body: " + err.Error()})
		return
	}

	var githubUser struct {
		ID        int64  `json:"id"`
		Login     string `json:"login"`
		Email     string `json:"email"`
		AvatarURL string `json:"avatar_url"`
	}

	if err := json.Unmarshal(data, &githubUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user info: " + err.Error()})
		return
	}

	var user User
	result := DB.Where("github_id = ?", githubUser.ID).First(&user)
	if result.Error != nil {
		user = User{
			GitHubID:       githubUser.ID,
			Username:       githubUser.Login,
			Email:          githubUser.Email,
			ProfilePicture: githubUser.AvatarURL,
			AccessToken:    token.AccessToken,
		}
		if err := DB.Create(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user: " + err.Error()})
			return
		}
	} else {
		user.Username = githubUser.Login
		user.Email = githubUser.Email
		user.ProfilePicture = githubUser.AvatarURL
		user.AccessToken = token.AccessToken
		DB.Save(&user)
	}

	jwtToken, err := GenerateJWT(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully authenticated with GitHub",
		"token": jwtToken,
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
		},
	})
}
