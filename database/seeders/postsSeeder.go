package seeders

import (
	"time"

	"github.com/Tej-11/connect-backend-application/database/config"
	"github.com/Tej-11/connect-backend-application/database/models"
)

func InsertSeedsInPosts() {
	var posts = []models.Post{
		{
			PostId: "1", 
			Caption: "Its Adhitya post",
			LikesCount: 150, 
			UserId: "1", 
			CreatedAt: time.Now(),
		},
		{
			PostId: "2", 
			Caption: "Its Karthik post", 
			LikesCount: 150, 
			UserId: "2", 
			CreatedAt: time.Now(),
		},
		{
			PostId: "3", 
			Caption: "Its Vyashu post", 
			LikesCount: 150, 
			UserId: "3", CreatedAt: time.Now(),
		},
		{
			PostId: "4", 
			Caption: "Its Suvega post", 
			LikesCount: 150, 
			UserId: "4", 
			CreatedAt: time.Now(),
		},
		{
			PostId: "5", 
			Caption: "Its Poorna post", 
			LikesCount: 150, 
			UserId: "5", 
			CreatedAt: time.Now(),
		},
		{
			PostId: "6", 
			Caption: "Its Poorna's second post", 
			LikesCount: 150, 
			UserId: "5", 
			CreatedAt: time.Now(),
		},
	}
	config.DB.Create(&posts)
}
