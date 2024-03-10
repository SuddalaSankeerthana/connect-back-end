package controllers

import (
	"net/http"
	"strings"

	"github.com/Tej-11/connect-backend-application/database/queries"
	"github.com/gin-gonic/gin"
)

func UpdateLikes(context *gin.Context) {

	postId := context.Query("postId")
	likeStatus := context.Query("likeStatus")

	postId = strings.ToLower(postId)
	likeStatus = strings.ToLower(likeStatus)

	if postId != "" && likeStatus != "" {

		if likeStatus == "true" || likeStatus == "false" {

			// if queries.IsPostPresent(postId) {

			var likesCount = queries.GetLikes(postId)

			if likeStatus == "true" {
				queries.UpdatePostLikes(postId, likesCount+1)
				likesCount = queries.GetLikes(postId)
				context.JSON(http.StatusOK, gin.H{
					"PostId":    postId,
					"LikeCount": likesCount,
				})

			} else {
				queries.UpdatePostLikes(postId, likesCount-1)
				likesCount = queries.GetLikes(postId)
				context.JSON(http.StatusOK, gin.H{
					"PostId":    postId,
					"LikeCount": likesCount,
				})

			}

			// } else {
			// 	context.JSON(http.StatusNotFound, gin.H{
			// 		"Error": "Provided post is not found",
			// 	})
			// }
		} else {
			context.JSON(http.StatusBadRequest, gin.H{
				"Error": "Like status cannot be other than \"true\" or \"false\" ",
			})
		}

	} else {
		context.JSON(http.StatusBadRequest, gin.H{
			"Error": "Please check if all query parameters are passed correctly",
		})
	}

}
