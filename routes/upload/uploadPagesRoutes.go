package upload

import (
	"github.com/Tej-11/connect-backend-application/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(routes *gin.Engine) {
	upload := routes.Group("/upload")

	upload.POST("/upload-post", controllers.UploadPost)
}
