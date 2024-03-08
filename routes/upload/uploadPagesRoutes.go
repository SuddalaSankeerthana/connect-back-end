package upload

import (
	"github.com/gin-gonic/gin"
)

func Routes(routes *gin.Engine) {
	upload := routes.Group("/upload")
}
