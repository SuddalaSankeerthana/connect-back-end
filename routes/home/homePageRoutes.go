package home

import (
	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	home := route.Group("/homepage")

}
