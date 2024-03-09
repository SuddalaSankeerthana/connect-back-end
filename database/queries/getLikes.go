package queries

import (
	"github.com/Tej-11/connect-backend-application/database/config"
	"github.com/Tej-11/connect-backend-application/database/models"
)

func GetLikes(postId string) int {

	var likesCount int

	config.DB.Model(&models.Post{}).
		Select("likes_count").
		Where("post_id =?", postId).Scan(&likesCount)
	return likesCount
}
