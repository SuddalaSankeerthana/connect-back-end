package customTypes

import "time"

type RawCommentsType struct {
	PostId              string
	UserId              string
	CommentId           string
	ParentCommentId     string
	Username            string
	ProfileImageAddress string
	CommentMessage      string
	CreatedAt           time.Time
}
