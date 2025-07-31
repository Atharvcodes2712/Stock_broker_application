package main

import (
	"authentication/router"
	"authentication/utils"
	"authentication/utils/db"
	"github.com/gin-gonic/gin"
)

func main() {
	utils.LoadConfig()
	db.InitDB()

	r := gin.Default()
	router.SetUpRoutes(r)

	utils.LogInfo("Server starting on port 8080")

	if err := r.Run(":8080"); err != nil {
		utils.LogError("Failed to start server: " + err.Error())
	}
}
