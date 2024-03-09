package customTypes

import "time"

type CommentType struct {
	CommentId           string
	UserId              string
	PostId              string
	Username            string
	ProfileImageAddress string
	CommentMessage      string
	CreatedAt           time.Time
	Replays             []ReplayType
}
