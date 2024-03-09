package home

import (
	"github.com/Tej-11/connect-backend-application/controllers"
	// "github.com/Tej-11/connect-backend-application/middleware"
	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	home := route.Group("/homepage")

	// home.GET("/get-posts", middleware.RequireAuth, controllers.GetPosts)
	// home.GET("/get-users", middleware.RequireAuth, controllers.GetUsers)
	home.GET("/get-posts", controllers.GetPosts)
	home.GET("/get-users", controllers.GetUsers)
	// home.GET("/get-posts", controllers.GetPosts)
	home.GET("/get-comments", controllers.GetComments)
	home.POST("/update-like", controllers.UpdateLikeTable)
	home.PATCH("/like", controllers.UpdateLikes)
}
