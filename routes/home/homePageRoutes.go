package home

import (
	"github.com/Tej-11/connect-backend-application/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	home := route.Group("/homepage")
	home.GET("/get-posts", controllers.GetPosts)
	home.GET("/get-comments", controllers.GetComments)
	home.POST("/update-like", controllers.UpdateLikeTable)
	home.PATCH("/like", controllers.UpdateLikes)
}
