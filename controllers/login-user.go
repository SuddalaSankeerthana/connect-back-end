package controllers

import (
	"net/http"
	"time"

	"github.com/Tej-11/connect-backend-application/database/config"
	"github.com/Tej-11/connect-backend-application/database/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("your_secret_key")

func LoginUser(c *gin.Context) {

	var loginData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := config.DB.Where("email = ?", loginData.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Email"})
		return
	}
	
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Password"})
		return
	}

	tokenString, err := generateToken(user.UserId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600 * 24 * 30, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{})

	// c.JSON(http.StatusOK, gin.H{"token": token, "user": user})
}

func generateToken(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userID,
		"exp":    time.Now().Add(time.Hour * 1).Unix(),
	})
	
	return token.SignedString(jwtKey)
}