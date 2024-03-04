package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	auth := route.Group("/auth")
    auth.POST("/login", func(c *gin.Context) {

		var requestBody struct {
			Message string `json:"message"`
		}

		if err := c.BindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "You are in the login route"})
	})
	auth.POST("/sign-up", func(c *gin.Context) {

		var requestBody struct {
			Message string `json:"message"`
		}

		if err := c.BindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "You are in the sign-up route"})
	})
}
