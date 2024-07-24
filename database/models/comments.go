package models

import "time"

type Comment struct {
	CommentId       string    `json:"commentId" gorm:"primary_key;type:varchar(255)"`
	CommentMessage  string    `json:"caption" gorm:"not null;type:varchar(255)"`
	UserId          string    `json:"userId" gorm:"not null;type:varchar(255)"`
	PostId          string    `json:"postId" gorm:"not null;type:varchar(255)"`
	ParentCommentId string    `json:"parentCommentId"`
	CreatedAt       time.Time `json:"CreatedAt"`
	Comments        []Comment `gorm:"foreignKey:ParentCommentId;references:CommentId;constraint:OnDelete:CASCADE;;type:varchar(255)"`
}
