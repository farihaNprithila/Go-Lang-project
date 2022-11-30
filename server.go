package main

import (
	"gitlab.com/pragmaticreviews/golang-gin-poc/src/middlewares"
	"gitlab.com/pragmaticreviews/golang-gin-poc/src/routes"
)

func main() {

	middlewares.SetupLogOutput()

	routes.Routes()
}
