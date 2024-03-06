# connect-backend-application

## Note
Create an .env file with the following data
- DSN = "user_name:password@tcp(host_name:port_number)/ConnectDevDB?charset=utf8mb4&parseTime=True&loc=Local"

## Installing packages
- To install all the dependencies/packages required use the command `go get ./...`

## Running the application
- To start the app use the command `go run main.go` or `go run .`
- Here we have our backend running in the address `http://localhost:8080/` in the following routes.
> - `/homepage/get-posts` -> This route fetches the posts data.
> - `/homepage/get-comments/:postId` -> This route fetches the comments data taking post id as query parameter.
> - `/homepage/like/:postId/:likeStatus` -> This route updates the likes count of a post taking post id and like status as query parameter.
>>- likeStatus value can be `liked` or `unliked` only

- I have also created the authentication routes:
> - POST `http://localhost:8080/auth/login` 
> - POST `http://localhost:8080/auth/sign-up`
>>> - For above two route takes body as `{
    "message": "Hello, world!"
}`
### > - While working with the /auth group change the functionality accordingly.

##
## Thank you.