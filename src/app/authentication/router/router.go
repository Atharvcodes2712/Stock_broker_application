package router

import (
	"authentication/constants"
	"authentication/handlers"
	"github.com/gin-gonic/gin"
)

func SetUpRoutes(router *gin.Engine) {
	router.POST(constants.SignInRoute, handlers.SignInHandler)
	router.POST(constants.SignUpRoute, handlers.SignUpHandler)
}
