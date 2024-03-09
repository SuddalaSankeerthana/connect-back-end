package queries

import (
	"github.com/Tej-11/connect-backend-application/database/config"
	"github.com/Tej-11/connect-backend-application/database/models"
)

func CreateImages(newPostImagesURLs map[string]string, newPostId string) error {

	for imageId, imageURL := range newPostImagesURLs {
		var image = models.Image{
			ImageId:  imageId,
			PostId:   newPostId,
			ImageURL: imageURL,
		}
		result := config.DB.Create(&image)

		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}
