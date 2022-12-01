package main

import (
	"gitlab.com/pragmaticreviews/golang-gin-poc/config"
	"gitlab.com/pragmaticreviews/golang-gin-poc/src/middlewares"
	"gitlab.com/pragmaticreviews/golang-gin-poc/src/routes"
	"gorm.io/gorm"
)

var db *gorm.DB = config.SetupDatabaseConnection()

func main() {

	middlewares.SetupLogOutput()
	defer config.CloseDatabaseConnection(db)

	routes.Routes()
}
