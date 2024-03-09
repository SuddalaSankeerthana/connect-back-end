package controllers

import (
	"github.com/Tej-11/connect-backend-application/database/queries"
	"github.com/Tej-11/connect-backend-application/utils"
	"github.com/gin-gonic/gin"
)

func GetPosts(context *gin.Context) {
	var postsRawData = queries.GetPosts(context)
	var posts = utils.QueryToPostsParser(postsRawData)
	context.JSON(200, gin.H{
		"posts": posts,
	})
}
