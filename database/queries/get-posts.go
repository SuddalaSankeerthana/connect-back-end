package queries

import (
	"fmt"

	"github.com/Tej-11/connect-backend-application/customTypes"
	"github.com/Tej-11/connect-backend-application/database/config"
	"github.com/Tej-11/connect-backend-application/database/models"
)

func GetPosts(userID string) []customTypes.ImagePostUser {

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
	userId := 
	const likedData = GetLikeStatus(userId)
	return queryData
}

func GetLikeStatus(userId string) {
	var queryData []models.Like
	config.DB.Table("likes").Select(
		`likes.post_id,likes.user_id`).Where(`likes.user_id`, userId).Scan(&queryData)
	fmt.Println(queryData)
	return
}
