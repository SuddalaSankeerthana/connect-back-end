package main

import (
	"github.com/Tej-11/connect-backend-application/database/config"
	"github.com/Tej-11/connect-backend-application/database/queries"
	"github.com/Tej-11/connect-backend-application/routes/auth"
	"github.com/Tej-11/connect-backend-application/routes/home"

	"github.com/gin-gonic/gin"
)

func main() {
	config.EstablishDatabaseConnection()
	queries.CreateSeeds()
	router := gin.Default()
	home.Routes(router)
	auth.Routes(router)
	router.Run()
}
