package main

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	GitHubID       int64  `json:"github_id"`
	Username       string `json:"username"`
	Email          string `json:"email"`
	ProfilePicture string `json:"profile_picture"`
	AccessToken    string `json:"-"`
}
