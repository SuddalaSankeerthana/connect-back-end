package seeders

import (
	"github.com/Tej-11/connect-backend-application/database/config"
	"github.com/Tej-11/connect-backend-application/database/models"
)

func InsertSeedsLikes() {
	var likes = []models.Like{
		{
			PostId: "1",
			UserId: "2",
		},
		{
			PostId: "2",
			UserId: "3",
		},
		{
			PostId: "3",
			UserId: "4",
		},
		{
			PostId: "4",
			UserId: "5",
		},
		{
			PostId: "5",
			UserId: "1",
		},
		{
			PostId: "6",
			UserId: "2",
		},
	}
	config.DB.Create(&likes)
}
