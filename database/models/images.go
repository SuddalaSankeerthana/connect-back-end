package models


type Image struct {
	ImageId   string    `json:"imageId" gorm:"primary_key;type:varchar(255)"`
	ImageURL  string    `json:"caption" gorm:"type:varchar(255)"`
	PostId    string    `json:"postId" gorm:"type:varchar(255);not null;"`
}
