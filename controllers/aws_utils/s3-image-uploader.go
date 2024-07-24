package aws_utils

import (
	"bytes"
	"fmt"
	"log"

	"github.com/Tej-11/connect-backend-application/env"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func S3UImagesUploader(images map[string][]byte, userId string, postId string) (map[string]string, error) {

	uploadedImagesURLs := make(map[string]string)

	accessKey := env.GetDotEnvVariable("AWS_ACCESS_KEY_ID")
	secertKey := env.GetDotEnvVariable("AWS_SECRET_ACCESS_KEY")
	bucketName := env.GetDotEnvVariable("BUCKET_NAME")

	aws_session, err := session.NewSession(&aws.Config{
		Region:      aws.String("ap-south-1"),
		Credentials: credentials.NewStaticCredentials(accessKey, secertKey, ""),
	})

	if err != nil {
		log.Fatal("Failed to create session", err)
	}

	s3BucketService := s3.New(aws_session)

	if postId != "" {

		for imageId, decodedImage := range images {

			s3BucketImagePath := "/" + userId + "/" + postId + "/" + imageId

			_, err = s3BucketService.PutObject(&s3.PutObjectInput{
				Bucket: aws.String(bucketName),
				Key:    aws.String(userId + "/" + postId + "/" + imageId),
				Body:   bytes.NewReader(decodedImage),
			})

			if err != nil {
				log.Fatal("Failed to upload data to S3", err)
				return uploadedImagesURLs, err

			} else {
				uploadedImagesURLs[imageId] = env.GetDotEnvVariable("S3BUCKET_URL") + s3BucketImagePath
			}
		}
	} else {
		for imageId, decodedImage := range images {

			s3BucketImagePath := "/" + userId

			_, err = s3BucketService.PutObject(&s3.PutObjectInput{
				Bucket: aws.String(bucketName),
				Key:    aws.String(userId),
				Body:   bytes.NewReader(decodedImage),
			})

			if err != nil {
				log.Fatal("Failed to upload data to S3", err)
				return uploadedImagesURLs, err

			} else {
				uploadedImagesURLs[userId] = env.GetDotEnvVariable("S3BUCKET_URL") + s3BucketImagePath
				fmt.Println(" ", uploadedImagesURLs, "", imageId)
			}
		}
	}

	return uploadedImagesURLs, nil
}
