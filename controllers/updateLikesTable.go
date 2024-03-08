package controllers

import (
	"net/http"

	"github.com/Tej-11/connect-backend-application/database/models"
	"github.com/Tej-11/connect-backend-application/database/queries"
	"github.com/gin-gonic/gin"
)

func UpdateLikeTable(context *gin.Context) {
	queryParams := context.Request.URL.Query()
	status := queryParams["status"][0]

	var likeData models.Like

	if err := context.BindJSON(&likeData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Updating likes table"})
	if status == "true" {
		queries.InsertLikes(likeData.PostId, likeData.UserId)
		context.JSON(http.StatusCreated, gin.H{"message": "Inserted data into likes table"})
	}
	if status == "false" {
		queries.DeleteLikes(likeData.PostId, likeData.UserId)
		context.JSON(http.StatusNoContent, gin.H{"message": "Delted row from likes table"})
	}
}
