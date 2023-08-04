package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wepull/bugTracker/models"
)

// RegisterUser is the handler for user registration.
func RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	// Validate the user data, check for duplicate usernames, etc.
	// Call the UserService to create the user and store it in the database.

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// LoginUser is the handler for user login.
func LoginUser(c *gin.Context) {
	var credentials struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	// Validate the user credentials.
	// Call the UserService to authenticate the user.

	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}
