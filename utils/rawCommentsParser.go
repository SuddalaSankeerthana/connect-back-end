package utils

import (
    "github.com/Tej-11/connect-backend-application/customTypes"
)

func RawCommentParser(rawCommentsData []customTypes.RawCommentsType) map[string]customTypes.CommentType {

    comments := make(map[string]customTypes.CommentType)

    for i := 0; i < len(rawCommentsData); i++ {

        var currentComment = rawCommentsData[i]

        if currentComment.ParentCommentId == "" {
            var newComment customTypes.CommentType
            var replays []customTypes.ReplayType
            newComment.PostId = currentComment.PostId
            newComment.UserId = currentComment.UserId
            newComment.CommentId = currentComment.CommentId
            newComment.CommentMessage = currentComment.CommentMessage
            newComment.ProfileImageAddress = currentComment.ProfileImageAddress
            newComment.Username = currentComment.Username
            newComment.CreatedAt = currentComment.CreatedAt
            newComment.Replays = replays

            comments[newComment.CommentId] = newComment
        }
    }

    for i := 0; i < len(rawCommentsData); i++ {

        var currentComment = rawCommentsData[i]

        if currentComment.ParentCommentId != "" {

            var newReplay customTypes.ReplayType

            newReplay.PostId = currentComment.PostId
            newReplay.UserId = currentComment.UserId
            newReplay.CommentId = currentComment.CommentId
            newReplay.ParentCommentId = currentComment.ParentCommentId
            newReplay.CommentMessage = currentComment.CommentMessage
            newReplay.ProfileImageAddress = currentComment.ProfileImageAddress
            newReplay.Username = currentComment.Username
            newReplay.CreatedAt = currentComment.CreatedAt

            val, ok := comments[currentComment.ParentCommentId]

            if ok {
                val.Replays = append(val.Replays, newReplay)
                comments[newReplay.ParentCommentId] = val
            }

        }
    }

    return comments
}