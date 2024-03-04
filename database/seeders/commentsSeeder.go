package seeders

import (
	"time"

	"github.com/Tej-11/connect-backend-application/database/config"
	"github.com/Tej-11/connect-backend-application/database/models"
)

func InsertSeedsInComments() {
	var comments = []models.Comment{
		{
			CommentId:      "1",
			CommentMessage: "Its parent comment 1",
			CreatedAt:      time.Now(),
			UserId:         "2",
			PostId:         "1",
		},
		{
			CommentId:      "2",
			CommentMessage: "Its parent comment 2",
			CreatedAt:      time.Now(),
			UserId:         "3",
			PostId:         "2",
		},
		{
			CommentId:      "3",
			CommentMessage: "Its parent comment 3",
			CreatedAt:      time.Now(),
			UserId:         "4",
			PostId:         "3",
		},
		{
			CommentId:      "4",
			CommentMessage: "Its parent comment 4",
			CreatedAt:      time.Now(),
			UserId:         "5",
			PostId:         "4",
		},
		{
			CommentId:      "5",
			CommentMessage: "Its parent comment 5",
			CreatedAt:      time.Now(),
			UserId:         "1",
			PostId:         "5",
		},
		{
			CommentId:      "6",
			CommentMessage: "Its parent comment 6",
			CreatedAt:      time.Now(),
			UserId:         "2",
			PostId:         "6",
		},

		{
			CommentId:       "7",
			CommentMessage:  "Its child comment 1",
			CreatedAt:       time.Now(),
			UserId:          "3",
			PostId:          "1",
			ParentCommentId: "1",
		},
		{
			CommentId:       "8",
			CommentMessage:  "Its child comment 2",
			CreatedAt:       time.Now(),
			UserId:          "4",
			PostId:          "2",
			ParentCommentId: "2",
		},
		{
			CommentId:       "9",
			CommentMessage:  "Its child comment 3",
			CreatedAt:       time.Now(),
			UserId:          "5",
			PostId:          "3",
			ParentCommentId: "3",
		},
		{
			CommentId:       "10",
			CommentMessage:  "Its child comment 4",
			CreatedAt:       time.Now(),
			UserId:          "1",
			PostId:          "4",
			ParentCommentId: "4",
		},
		{
			CommentId:       "11",
			CommentMessage:  "Its child comment 5",
			CreatedAt:       time.Now(),
			UserId:          "2",
			PostId:          "5",
			ParentCommentId: "5",
		},
		{
			CommentId:       "12",
			CommentMessage:  "Its child comment 6",
			CreatedAt:       time.Now(),
			UserId:          "3",
			PostId:          "6",
			ParentCommentId: "6",
		},
	}
	config.DB.Create(&comments)
}
