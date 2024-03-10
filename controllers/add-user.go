package controllers

import (
	"net/http"

	"github.com/Tej-11/connect-backend-application/controllers/aws_utils"
	"github.com/Tej-11/connect-backend-application/customTypes"
	"github.com/Tej-11/connect-backend-application/database/models"
	"github.com/Tej-11/connect-backend-application/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func AddUsers(c *gin.Context) {
	var user customTypes.RegisteredDetailsBodyType
	var newUser models.User

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

	imagesURLs, err := aws_utils.S3UImagesUploader(base64DecodedImages, user.UserId, "")

	tokenString, err := generateToken(user.UserId)

	newUser.UserId = user.UserId
	newUser.Username = user.Username
	newUser.Email = user.Email
	newUser.Password = user.Password
	newUser.ProfileImageAddress = imagesURLs[user.UserId]

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	// c.SetSameSite(http.SameSiteLaxMode)
	// c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully", "user": user, "token": tokenString})

}
