package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/wepull/bugTracker/models"
	"github.com/wepull/bugTracker/utils"
)

// RegisterUser is the handler for user registration.
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	resp := models.ApiResp{
		ResponseCode:        http.StatusCreated,
		ResponseDescription: "User registered successfully",
	}

	defer func() {
		out, _ := json.Marshal(resp)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Add("Content-Type", "application/json")
		_, err := w.Write(out)
		if err != nil {
			fmt.Printf("%v", err)
		}
	}()
	var req models.RegisterUserRequest
	if err := utils.ReadAndParseInput(w, r, &req); err != nil {
		resp.ResponseCode = http.StatusBadRequest
		resp.ResponseDescription = err.Error()
		return
	}

	if !utils.IsValidEmail(req.Email) {
		resp.ResponseCode = http.StatusBadRequest
		resp.ResponseDescription = "Error Registering user, please provide a valid Email"
	}
}

// LoginUser is the handler for user login.
func LoginUser(w http.ResponseWriter, r *http.Request) {
	// var credentials struct {
	// 	Username string `json:"username" binding:"required"`
	// 	Password string `json:"password" binding:"required"`
	// }
	// if err := c.ShouldBindJSON(&credentials); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
	// 	return
	// }

	// // Validate the user credentials.
	// // Call the UserService to authenticate the user.

	// c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}
