package queries

import (
	"time"

	"github.com/Tej-11/connect-backend-application/customTypes"
	"github.com/Tej-11/connect-backend-application/database/config"
	"github.com/Tej-11/connect-backend-application/database/models"

	"github.com/google/uuid"
)

func CrearteComment(commentDetails customTypes.NewCommentBodyType) bool {

	var newCommentData models.Comment
	newCommentData.PostId = commentDetails.PostId
	newCommentData.UserId = commentDetails.UserId
	newCommentData.CommentId = uuid.New().String()
	newCommentData.CommentMessage = commentDetails.CommentMessage
	newCommentData.ParentCommentId = commentDetails.ParentCommentId
	newCommentData.CreatedAt = time.Now()

	result := config.DB.Create(&newCommentData)

	if result.Error != nil {
		return true
	} else {
		return false
	}

}
