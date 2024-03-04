package models

type User struct {
	UserId              string    `json:"UserId" gorm:"primaryKey;type:varchar(255)"`
	Username            string    `json:"Username" gorm:"unique;not null;type:varchar(255)"`
	Email               string    `json:"Email" gorm:"unique;not null;type:varchar(255)"`
	Password            string    `json:"Password" gorm:"not null;type:varchar(255)"`
	ProfileImageAddress string    `json:"ProfileImageAddress" gorm:"type:varchar(255)"`
	Posts               []Post    `gorm:"foreignKey:UserId;references:UserId;constraint:OnDelete:CASCADE;"`
	Comments            []Comment `gorm:"foreignKey:UserId;references:UserId;constraint:OnDelete:CASCADE;"`
	Likes               []Like    `gorm:"foreignKey:UserId;references:UserId;constraint:OnDelete:CASCADE;"`
}
