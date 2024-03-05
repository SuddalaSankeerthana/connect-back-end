package queries

import (
	"fmt"

	"github.com/Tej-11/connect-backend-application/database/config"
	"github.com/Tej-11/connect-backend-application/database/models"
	"gorm.io/gorm"
)

func IsPostPresent(postId string) bool {

	var post models.Post
	id := postId
	if err := config.DB.First(&post, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false
		} else {
			return false
			fmt.Println("Error:", err)
		}
	}
	return true

}
