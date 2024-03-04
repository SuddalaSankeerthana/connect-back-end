package models

type Like struct {
	UserId string `json:"UserId" gorm:"primaryKey;type:varchar(255)"`
	PostId string `json:"PostId" gorm:"primaryKey;type:varchar(255)"`
}
