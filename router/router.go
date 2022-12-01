package router

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/pragmaticreviews/golang-gin-poc/controller"
	"gitlab.com/pragmaticreviews/golang-gin-poc/controller/authentication"
	"gitlab.com/pragmaticreviews/golang-gin-poc/middlewares"
	"gitlab.com/pragmaticreviews/golang-gin-poc/service"
	"net/http"
)

var (
	authenticationController authentication.AuthController = authentication.NewAuthController()
	userService              service.UserService           = service.New()
	userController           controller.UserController     = controller.New(userService)
)

func Router() {
	route := gin.New()

	route.Use(gin.Recovery(), middlewares.Logger())

	apiRoutes := route.Group("/api")
	{
		authRoutes := apiRoutes.Group("/auth")
		{
			authRoutes.POST("/login", authenticationController.Login)
			authRoutes.POST("/register", authenticationController.Register)
		}
		apiRoutes.GET("/users", func(ctx *gin.Context) {
			ctx.JSON(200, userController.FindAll())

		})
		apiRoutes.POST("/users", func(ctx *gin.Context) {
			err := userController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "User input was valid"})
			}
		})
	}

	viewRoutes := route.Group("/view")
	{
		viewRoutes.GET("/users", userController.ShowAll)
	}

	err := route.Run(":8080")
	if err != nil {
		return
	}
}
