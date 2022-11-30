package routes

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/pragmaticreviews/golang-gin-poc/src/controller"
	"gitlab.com/pragmaticreviews/golang-gin-poc/src/middlewares"
	"gitlab.com/pragmaticreviews/golang-gin-poc/src/service"
)

var (
	userService    service.UserService       = service.New()
	userController controller.UserController = controller.New(userService)
)

func Routes() {
	route := gin.New()

	route.Use(gin.Recovery(), middlewares.Logger())

	route.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(200, userController.FindAll())

	})
	route.POST("/users", func(ctx *gin.Context) {
		ctx.JSON(200, userController.Save(ctx))
	})

	err := route.Run(":8080")
	if err != nil {
		return
	}
}
