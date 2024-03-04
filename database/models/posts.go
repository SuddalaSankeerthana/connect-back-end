package models

import "time"

type Post struct {
	PostId     string    `json:"PostId" gorm:"primary_key;type:varchar(255)"`
	Caption    string    `json:"Caption" gorm:"type:varchar(255)"`
	LikesCount int       `json:"LikesCount" gorm:"default:0"`
	UserId     string    `json:"UserId" gorm:"type:varchar(255)"`
	CreatedAt  time.Time `json:"CreatedAt"`
	Images     []Image   `gorm:"foreignKey:PostId;references:PostId;constraint:OnDelete:CASCADE;"`
	Comments   []Comment `gorm:"foreignKey:PostId;references:PostId;constraint:OnDelete:CASCADE;"`
	Likes      []Like    `gorm:"foreignKey:PostId;references:PostId;constraint:OnDelete:CASCADE;"`
}
