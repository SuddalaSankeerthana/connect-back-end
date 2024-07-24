package customTypes

import "time"

type ReplayType struct {
	CommentId           string
	UserId              string
	PostId              string
	ParentCommentId     string
	Username            string
	ProfileImageAddress string
	CommentMessage      string
	CreatedAt           time.Time
}
