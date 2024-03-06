package auth

import (
	"fmt"
	"net/http"
	// "time"
	"github.com/gin-gonic/gin"
	// "github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"github.com/Tej-11/connect-backend-application/database/config"
	"github.com/Tej-11/connect-backend-application/database/models"
)


func Routes(route *gin.Engine) {

	auth := route.Group("/auth")
	auth.POST("/register", func(c *gin.Context) {
        var user models.User

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
			return
		}

		user.UserId = uuid.New().String()
		user.Password = string(hashedPassword)
		fmt.Println(string(hashedPassword))

		if err := config.DB.Create(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})

	})

}
