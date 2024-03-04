package seeders

import (
	"github.com/Tej-11/connect-backend-application/database/config"
	"github.com/Tej-11/connect-backend-application/database/models"
)

func InsertSeedsInUsers() {
	var users = []models.User{
		{
			UserId:   "1",
			Username: "Adhitya",
			Email:    "adhitya@gmail.com",
			Password: "$2a$14$uBCeus2RJ3RhxYnzC5wpdezPydAtGdctvJpvqvrK7Rr7kHRhifKge", 
			ProfileImageAddress: "https://static.vecteezy.com/system/resources/thumbnails/006/541/188/small/bearded-male-cartoon-character-with-sunglasses-minimalist-cartoon-avatar-profile-illustration-free-vector.jpg",
		},
		{
			UserId: "2", 
			Username: "Karthik", 
			Email: "karthik@gmail.com", 
			Password: "$2a$14$uBCeus2RJ3RhxYnzC5wpdezPydAtGdctvJpvqvrK7Rr7kHRhifKge", 
			ProfileImageAddress: "https://cdn.analyticsvidhya.com/wp-content/uploads/2023/08/Magic-AI-.png",
		},
		{
			UserId: "3", 
			Username: "Vyshu", 
			Email: "vyshu@gmail.com", 
			Password: "$2a$14$uBCeus2RJ3RhxYnzC5wpdezPydAtGdctvJpvqvrK7Rr7kHRhifKge", 
			ProfileImageAddress: "https://images.pexels.com/photos/415829/pexels-photo-415829.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=2",
		},
		{
			UserId: "4", 
			Username: "Suvega", 
			Email: "suvega@gmail.com", 
			Password: "$2a$14$uBCeus2RJ3RhxYnzC5wpdezPydAtGdctvJpvqvrK7Rr7kHRhifKge", 
			ProfileImageAddress: "https://png.pngtree.com/png-clipart/20230927/original/pngtree-man-avatar-image-for-profile-png-image_13001882.png",
		},
		{
			UserId: "5", 
			Username: "Poorna", 
			Email: "poorna@gmail.com", 
			Password: "$2a$14$uBCeus2RJ3RhxYnzC5wpdezPydAtGdctvJpvqvrK7Rr7kHRhifKge", 
			ProfileImageAddress: "https://media.istockphoto.com/id/1483487034/photo/portrait-of-a-cute-female-video-game-avatar.webp?b=1&s=170667a&w=0&k=20&c=tMwM3rBZ6eOhv9LQi2lTUgfbC96X2BDKAmpUpK2r-BA=",
		},
	}
	config.DB.Create(&users)
}
