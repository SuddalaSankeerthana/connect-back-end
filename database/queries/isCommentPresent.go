package queries

import (
	"github.com/Tej-11/connect-backend-application/database/config"
	"github.com/Tej-11/connect-backend-application/database/models"
	"gorm.io/gorm"
)

func IsCommentPresent(commentId string) bool {

	var comment models.Comment
	id := commentId
	if err := config.DB.First(&comment, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false
		} else {
			return false
		}
	}
	return true
}
