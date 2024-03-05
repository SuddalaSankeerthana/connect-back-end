package home

import (
	"github.com/Tej-11/connect-backend-application/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	home := route.Group("/homepage")
	home.GET("/get-posts", controllers.GetPosts)
	home.GET("/get-comments", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "You are on the get-comments route",
		})
	})
}
