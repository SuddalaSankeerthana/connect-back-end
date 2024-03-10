package controllers

import (
	"net/http"

	"github.com/Tej-11/connect-backend-application/customTypes"
	"github.com/Tej-11/connect-backend-application/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func AddUsers(c *gin.Context) {
	var user customTypes.RegisteredDetailsBodyType

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

	base64DecodedImages, err := utils.Base64Decoder(user.ProfileImageAddress)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode profile picture"})
		return
	}

	// c.SetSameSite(http.SameSiteLaxMode)
	// c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully", "user": user, "token": tokenString})

}
