package controllers

import (
	"fmt"

	"github.com/Tej-11/connect-backend-application/database/queries"
	"github.com/Tej-11/connect-backend-application/utils"
	"github.com/gin-gonic/gin"
)

func GetComments(context *gin.Context) {

	queryParams := context.Request.URL.Query()
	postId := queryParams["postId"][0]
	fmt.Println(postId)

	// if queries.IsPostPresent(postId) {
	var commentsRawData = queries.GetComments(postId)
	var comments = utils.RawCommentParser(commentsRawData)
	context.JSON(200, gin.H{
		"comments": comments,
	})
	// } else {
	// 	context.JSON(404, gin.H{
	// 		"Error": "Post doesn't exist",
	// 	})
	// }

}
