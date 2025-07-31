package handlers

import (
	"authentication/models"
	"authentication/service"
	"authentication/utils"
	"authentication/utils/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SignUpHandler(c *gin.Context) {
	var req models.SignUpRequest
	utils.LogInfo(" Received signup request")

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.LogError("Signup failed for user:" + req.UserName)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	msg, err := service.RegisterUser(req, db.DB)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	utils.LogInfo("Signup successful for user:" + req.UserName)
	c.JSON(http.StatusCreated, gin.H{"message": msg})
}
