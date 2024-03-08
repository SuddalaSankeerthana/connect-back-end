package upload

import (
	"github.com/gin-gonic/gin"
)

func Routes(routes *gin.Engine) {
	upload := routes.Group("/upload")

	upload.POST("/upload-post", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "This is upload route",
		})
	})
}
