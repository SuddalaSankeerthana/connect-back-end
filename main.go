package main

import (
	"github.com/Tej-11/connect-backend-application/database/config"
	"github.com/Tej-11/connect-backend-application/database/queries"
	"github.com/Tej-11/connect-backend-application/routes/auth"
	"github.com/Tej-11/connect-backend-application/routes/home"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// func GetLikeStatus(userId string) {
// 	var queryData []models.Like
// 	config.DB.Table("likes").Select(
// 		`likes.post_id,likes.user_id`).Where(`likes.user_id`, userId).Scan(&queryData)
// 	fmt.Println(queryData)
// 	return
// }

type LikeStatus struct {
	likeStatus string
	postId     string
}

func main() {
	config.EstablishDatabaseConnection()
	queries.GetLikeStatus("2")
	// queries.CreateSeeds()
	router := gin.Default()
	router.Use(cors.Default())
	home.Routes(router)
	auth.Routes(router)
	router.Run()
}
