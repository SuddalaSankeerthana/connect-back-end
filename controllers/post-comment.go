package controllers

import (
	"net/http"

	"github.com/Tej-11/connect-backend-application/customTypes"
	"github.com/Tej-11/connect-backend-application/database/queries"

	"github.com/gin-gonic/gin"
)

func PostComment(context *gin.Context) {
	var newCommentData customTypes.NewCommentBodyType
	userId := newCommentData.UserId
	postId := newCommentData.PostId
	parentCommentId := newCommentData.ParentCommentId
	context.BindJSON(&newCommentData)

	if newCommentData.UserId == "" || newCommentData.PostId == "" || newCommentData.CommentMessage == "" {

		if (newCommentData.UserId == "" || newCommentData.PostId == "") && newCommentData.PostId != "" {
			context.JSON(http.StatusNotFound, gin.H{"error": "Please provide valid data"})
			return
		}
		if newCommentData.CommentMessage == "" {
			context.JSON(http.StatusNotFound, gin.H{"error": "Comment cannot be empty"})
			return
		}
		context.JSON(http.StatusNotFound, gin.H{"error": "Please provide valid data"})
		return

	}

	if newCommentData.UserId != "" && newCommentData.PostId != "" && newCommentData.CommentMessage != "" {

		if !queries.IsPostPresent(postId) {
			context.JSON(http.StatusNotFound, gin.H{"error": "Invalid post alert"})
			return
		}
		if !queries.IsUserPresent(userId) {
			context.JSON(http.StatusNotFound, gin.H{"error": "Invalid user alert"})
			return
		}
		if(newCommentData.ParentCommentId!="" && queries.IsCommentPresent(parentCommentId)){

			context.JSON(http.StatusNotFound, gin.H{"error": "Please provide valid data 8"})
			return
		}

		queryResult := queries.CrearteComment(newCommentData)

		if !queryResult {
			context.JSON(http.StatusOK, gin.H{"message": "Comment uploaded successfully"})
			return
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"new comment": "Internal Server Error :("})
			return
		}
	}

	context.JSON(http.StatusInternalServerError, gin.H{"new comment": "Internal Server Error :("})
}
