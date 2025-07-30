package main

import (
	"github.com/gin-gonic/gin"
	"authentication/router"
	"authentication/utils"
	"authentication/utils/db"
)

func main() {
	utils.LoadConfig()
	db.InitDB()

	// Set up Gin router
	r := gin.Default()
	router.SetupRoutes(r)

	utils.LogInfo("Server starting on port 8080")

	// Start server
	if err := r.Run(":8080"); err != nil {
		utils.LogError("Failed to start server: " + err.Error())
	}
}
