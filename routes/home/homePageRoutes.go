package home

import (
	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	home := route.Group("/homepage")
	home.GET("/get-posts", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "All posts",
		})
	})
}
