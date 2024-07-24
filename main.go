package main

import (
	"github.com/Tej-11/connect-backend-application/database/config"
	"github.com/Tej-11/connect-backend-application/database/queries"
	// "github.com/Tej-11/connect-backend-application/database/queries"
	"github.com/Tej-11/connect-backend-application/routes/auth"
	"github.com/Tej-11/connect-backend-application/routes/home"
	"github.com/Tej-11/connect-backend-application/routes/upload"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.EstablishDatabaseConnection()
	queries.CreateSeeds()
	router := gin.Default()

	// configure := cors.DefaultConfig()
	// configure.AllowOrigins = []string{"http://localhost:3000"}
	// configure.AllowCredentials = true
	// configure.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Cookie"}

	// router.Use(cors.New(configure))
	router.Use(cors.Default())
	home.Routes(router)
	auth.Routes(router)
	upload.Routes(router)
	router.Run()
}
