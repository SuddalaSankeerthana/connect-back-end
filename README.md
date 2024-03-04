# connect-backend-application

## Note
Create an .env file with the following data
- DSN = "user_name:password@tcp(localhost:3306)/ConnectDevDB?charset=utf8mb4&parseTime=True&loc=Local"

## Installing packages
- To install all the dependencies/packages required use the command `go get ./...`

## Running the application
- To start the app use the command `go run main.go` or `go run .`
- Here we have our backend running in the address `http://localhost:8080/homepage/get-posts`. This route fetches the posts data.
- I have also created the authentication routes:
> - POST `http://localhost:8080/auth/login` 
> - POST `http://localhost:8080/auth/sign-up`
>> - For above two route takes body as `{
    "message": "Hello, world!"
}`
### > - While working with the /auth group change the functionality accordingly.

##
## Thank you.