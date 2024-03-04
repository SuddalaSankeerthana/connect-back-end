package config

import (
	"fmt"

	"github.com/Tej-11/connect-backend-application/env"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func EstablishDatabaseConnection() (*gorm.DB, error) {

	dsn := env.GetDotEnvVariable("DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	} else {
		fmt.Println("connection established")
	}

	DB = db

	return db, nil
}
