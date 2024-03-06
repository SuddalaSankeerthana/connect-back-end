package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/Tej-11/connect-backend-application/controllers"
)


func Routes(route *gin.Engine) {

	auth := route.Group("/auth")
	auth.POST("/register", controllers.AddUsers)
    auth.POST("/login", controllers.LoginUser)
}
