package router

import (
	"github.com/gin-gonic/gin"
	"authentication/constants"
	"authentication/handlers"
)

func SetupRoutes(r *gin.Engine) {
	r.POST(constants.SignUpRoute, handlers.SignUp)
	r.POST(constants.SignInRoute, handlers.SignIn)
}
