package queries

import (
	"time"

	"github.com/Tej-11/connect-backend-application/database/config"
	"github.com/Tej-11/connect-backend-application/database/models"
)

func CreatePost(userId string, newPostId string, caption string) error {

	var posts = []models.Post{
		{
			PostId:    newPostId,
			Caption:   caption,
			UserId:    userId,
			CreatedAt: time.Now(),
		},
	}

	result := config.DB.Create(&posts)
	
	if result.Error != nil {
		return result.Error
	}

	return nil
}
