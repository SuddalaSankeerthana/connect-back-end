package queries

import (
	"github.com/Tej-11/connect-backend-application/database/config"
	"github.com/Tej-11/connect-backend-application/database/models"
)

func GetLikesTableData(userId string) []models.Like {
	var queryData []models.Like
	config.DB.Table("likes").Select(
		`likes.post_id,likes.user_id`).Where(`likes.user_id`, userId).Scan(&queryData)
	return queryData
}
