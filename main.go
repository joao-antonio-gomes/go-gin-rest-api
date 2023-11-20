package main

import (
	"github.com/joao-antonio-gomes/go-gin-rest-api/database"
	"github.com/joao-antonio-gomes/go-gin-rest-api/routes"
)

func main() {
	database.ConnectDatabase()

	routes.HandleRequests()
}
