package main

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/pragmaticreviews/golang-gin-poc/config"
	"gitlab.com/pragmaticreviews/golang-gin-poc/middlewares"
	"gitlab.com/pragmaticreviews/golang-gin-poc/router"
	"gorm.io/gorm"
)

var db *gorm.DB = config.SetupDatabaseConnection()

func main() {
	gin.ForceConsoleColor()

	middlewares.SetupLogOutput()
	defer config.CloseDatabaseConnection(db)

	router.Router()
}
