package utils

import (
	"github.com/Tej-11/connect-backend-application/customTypes"
)

func QueryToPostsParser(queryData []customTypes.ImagePostUser) map[string]customTypes.PostType {

	posts := make(map[string]customTypes.PostType)

	for i := 0; i < len(queryData); i++ {
		var data customTypes.ImagePostUser = queryData[i]
		val, ok := posts[data.PostId]

		if ok {
			var images []string = val.Images
			images = append(images, data.ImageURL)
			val.Images = images
			posts[data.PostId] = val
		} else {

			var postDetails customTypes.PostType
			postDetails.PostId = data.PostId
			postDetails.UserId = data.UserId
			postDetails.Username = data.Username
			postDetails.ProfileImageAddress = data.ProfileImageAddress
			postDetails.Images = []string{data.ImageURL}
			postDetails.Caption = data.Caption
			postDetails.LikesCount = data.LikesCount

			posts[data.PostId] = postDetails
		}
	}
	return posts

}
