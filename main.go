package main

import (
	"github.com/Tej-11/connect-backend-application/database/config"
	"github.com/Tej-11/connect-backend-application/database/queries"
)

func main() {
	config.EstablishDatabaseConnection()
	queries.CreateSeeds()
}
