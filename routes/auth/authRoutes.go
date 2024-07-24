package auth

import (
	"github.com/Tej-11/connect-backend-application/controllers"
	"github.com/gin-gonic/gin"
)


func Routes(route *gin.Engine) {

	auth := route.Group("/auth")
	auth.POST("/register", controllers.AddUsers)
    auth.POST("/login", controllers.LoginUser)
}
