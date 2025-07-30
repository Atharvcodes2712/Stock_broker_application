
package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"authentication/constants"
	"authentication/models"
	"authentication/repo"
	"authentication/utils"
)



func SignUp(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		utils.LogError("Invalid signup input: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input format"})
		return
	}

	if user.UserName == "" || user.Password == "" || user.Email == "" || user.PhoneNumber == "" || user.PanCard == "" {
		utils.LogError("Missing required signup fields")
		c.JSON(http.StatusBadRequest, gin.H{"error": "All fields are required"})
		return
	}

	err := repo.CreateUser(user)
	if err != nil {
		utils.LogError("Signup error: " + err.Error())

		switch err {
		case constants.ErrUserAlreadyExists:
			c.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
		case constants.ErrEmailAlreadyExists:
			c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
		case constants.ErrCreatingUser:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unknown error"})
		}
		return
	}

	utils.LogInfo("New user created: " + user.UserName)
	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}
func SignIn(c *gin.Context) {
	var login models.User

	if err := c.ShouldBindJSON(&login); err != nil {
		utils.LogError("Invalid login input: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input format"})
		return
	}


	valid, err := repo.ValidateUser(login.UserName, login.Password)
	if err != nil {
		utils.LogError("Login error: " + err.Error())
		switch err {
		case constants.ErrUserNotFound:
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		case constants.ErrInvalidCred:
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		}
		return
	}

	if !valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	utils.LogInfo("Login success: " + login.UserName)
	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

