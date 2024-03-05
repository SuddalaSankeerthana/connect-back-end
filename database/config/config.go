package config

import (
	"fmt"

	"github.com/Tej-11/connect-backend-application/database/models"
	"github.com/Tej-11/connect-backend-application/env"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func EstablishDatabaseConnection() (*gorm.DB, error) {

	dsn := env.GetDotEnvVariable("DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	} else {
		fmt.Println("connection established")
		db.AutoMigrate(&models.User{}, &models.Post{}, &models.Image{}, &models.Comment{}, &models.Like{})
	}

	DB = db

	return db, nil
}
