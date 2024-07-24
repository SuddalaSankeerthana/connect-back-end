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
func InsertLikes(postId string, userId string) {
	config.DB.Model(&models.Like{}).Create(&models.Like{PostId: postId, UserId: userId})
}
func DeleteLikes(postId string, userId string) {
	config.DB.Model(&models.Like{}).Delete(models.Like{PostId: postId, UserId: userId})
}
