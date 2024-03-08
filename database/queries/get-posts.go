package queries

import (
	"fmt"

	"github.com/Tej-11/connect-backend-application/customTypes"
	"github.com/Tej-11/connect-backend-application/database/config"
	"github.com/Tej-11/connect-backend-application/database/models"
	"github.com/gin-gonic/gin"
)

func GetPosts(context *gin.Context) []customTypes.ImagePostUser {
	queryParams := context.Request.URL.Query()
	for key, values := range queryParams {
		fmt.Printf("Parameter: %s, Values: %v\n", key, values)
	}
	var likeStatusData = GetLikeStatus(queryParams["userId"][0])
	var queryData []customTypes.ImagePostUser
	config.DB.Table("images").
		Select(`
    users.user_id, 
    users.username, 
    users.profile_image_address,
    posts.post_id, 
	posts.caption, 
    posts.likes_count, 
    images.image_url`).
		Joins("right join posts on posts.post_id = images.post_id").
		Joins("right join users on users.user_id = posts.user_id").
		Scan(&queryData)

	for postIndex := 0; postIndex < len(queryData); postIndex++ {
		for likeIndex := 0; likeIndex < len(likeStatusData); likeIndex++ {
			if queryData[postIndex].PostId == likeStatusData[likeIndex].PostId {
				queryData[postIndex].LikeStatus = true
			}
		}
	}
	return queryData
}

func GetLikeStatus(userId string) []models.Like {
	var queryData []models.Like
	config.DB.Table("likes").Select(
		`likes.post_id,likes.user_id`).Where(`likes.user_id`, userId).Scan(&queryData)
	return queryData
}
