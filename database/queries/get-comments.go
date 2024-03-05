package queries

import (
	"github.com/Tej-11/connect-backend-application/customTypes"
	"github.com/Tej-11/connect-backend-application/database/config"
)

func GetComments(postId string) []customTypes.RawCommentsType {

	var rawCommentsData []customTypes.RawCommentsType

	config.DB.Table("comments").
		Select(`
			comments.comment_id,
			comments.user_id, 
			comments.post_id,
			comments.parent_comment_id,
			users.username, 
			users.profile_image_address,
			comments.comment_message, 
			comments.created_at
			`).
		Where("comments.post_id = ?", postId).
		Joins("right join users on users.user_id = comments.user_id").
		Scan(&rawCommentsData)

	// comments := utils.RawCommentParser(rawCommentsData)

	return rawCommentsData
}
