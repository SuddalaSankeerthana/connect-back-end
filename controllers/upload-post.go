package controllers

import (
	"log"

	"github.com/Tej-11/connect-backend-application/controllers/aws_utils"
	"github.com/Tej-11/connect-backend-application/customTypes"
	"github.com/Tej-11/connect-backend-application/database/queries"
	"github.com/Tej-11/connect-backend-application/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UploadPost(context *gin.Context) {

	var userId string
	postId := uuid.NewString()
	var caption string
	var postBody customTypes.PostBodyData
	base64DecodedImages := make(map[string][]byte)
	var imagesURLs map[string]string

	if err := context.BindJSON(&postBody); err != nil {
		context.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	userId = postBody.UserId
	caption = postBody.PostCaption

	if len(postBody.Images) == 0 {
		log.Fatal("Images were not provided")
		context.JSON(400, gin.H{
			"Error": "Incorrect data in body provided",
		})
		return
	}

	base64DecodedImages, err := utils.Base64Decoder(postBody.Images)
	if err != nil {
		log.Fatal("Couldn't decode base64 images")
		context.JSON(500, gin.H{
			"Error": "Internal Server Error",
		})
	}

	imagesURLs, err = aws_utils.S3UImagesUploader(base64DecodedImages, userId, postId)

	if err != nil {
		log.Fatal("Couldn't upload the images to the S3 bucket")
		context.JSON(500, gin.H{
			"Error": "Internal Server Error",
		})
	}

	if len(imagesURLs) < 1 {
		log.Fatal("Couldn't upload posts to S3")
		context.JSON(500, gin.H{
			"Error": "Internal Server Error",
		})
	}
	err = queries.CreatePost(userId, postId, caption)
	if err != nil {
		log.Fatal("Couldn't insert post data into the posts table")
		context.JSON(500, gin.H{
			"Error": "Internal Server Error",
		})
	}

	err = queries.CreateImages(imagesURLs, postId)

	if err != nil {
		log.Fatal("Couldn't insert images data into the images table")
		context.JSON(500, gin.H{
			"Error": "Internal Server Error",
		})
	}

	context.JSON(200, gin.H{
		"Success": "Post uploaded successfully \n you are ready to CONNECT",
	})
}
