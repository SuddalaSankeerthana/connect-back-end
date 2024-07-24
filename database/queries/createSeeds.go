package queries

import "github.com/Tej-11/connect-backend-application/database/seeders"

func CreateSeeds() {
	seeders.InsertSeedsInUsers()
	seeders.InsertSeedsInPosts()
	seeders.InsertSeedsInImages()
	seeders.InsertSeedsInComments()
	seeders.InsertSeedsLikes()
}
