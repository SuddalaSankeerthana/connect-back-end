# connect-backend-application: Upload

This repository contains the upload posts of the backend application for Connect. It handles functionalities related to uploading posts and images.

## Environment Setup

- Create an .env file in the root directory with the following data:

```
DSN = "user_name:password@tcp(host_name:port_number)/ConnectDevDB?charset=utf8mb4&parseTime=True&loc=Local"
AWS_ACCESS_KEY_ID=your_aws_access_key_id
AWS_SECRET_ACCESS_KEY=your_aws_secret_access_key
BUCKET_NAME=your_s3_bucket_name
S3BUCKET_URL=your_s3_bucket_url

```

## Installing Packages

- To install all the required dependencies, use the following command:

` go get ./...`

## Running the Application

- To start the app use the command `go run main.go` or `go run .`

## Usage

### Endpoints

- /upload/upload-post [POST]
- This endpoint is used to upload a post with images to the server.

- _Request Body_:

```
{
  "UserId": "user_id",
  "PostCaption": "post_caption",
  "Images": {
    "image_id_1": "base64_encoded_image_data_1",
    "image_id_2": "base64_encoded_image_data_2",
    ...
  }
}
```

- _Response_:

1. Success: 200 OK

```
{
  "Success": "Post uploaded successfully. You are ready to CONNECT."
}
```

2. Error: 500 Internal Server Error

```
{
  "Error": "Internal Server Error"
}
```

## License

- This project is licensed under the MIT License.
