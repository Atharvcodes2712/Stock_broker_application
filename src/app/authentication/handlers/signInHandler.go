package handlers

import (
	"authentication/service"
	"authentication/utils"
	"authentication/utils/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func SignInHandler(c *gin.Context) {
	var req LoginRequest
	utils.LogInfo("sign in reqest received for user" + req.Username)
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.LogError("Invalid input during SignIn")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	msg, err := service.CheckCredentials(req.Username, req.Password, db.DB)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	utils.LogInfo("login successful")
	c.JSON(http.StatusOK, gin.H{"message": msg})
}
