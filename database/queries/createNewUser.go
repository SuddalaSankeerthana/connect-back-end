package queries

import (
	"github.com/Tej-11/connect-backend-application/database/config"
	"github.com/Tej-11/connect-backend-application/database/models"
)

func CreateNewUser(newUser models.User) error {

	result := config.DB.Create(&newUser)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
