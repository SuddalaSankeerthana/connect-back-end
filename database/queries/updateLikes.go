package queries

import (
	"github.com/Tej-11/connect-backend-application/database/config"
	"github.com/Tej-11/connect-backend-application/database/models"
)

func UpdatePostLikes(postId string, likesCount int) {
	
	config.DB.Model(&models.Post{}).
		Where("post_id = ?", postId).
		Update("likes_count", likesCount)

}
