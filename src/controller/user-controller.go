package controller

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/pragmaticreviews/golang-gin-poc/src/entity"
	"gitlab.com/pragmaticreviews/golang-gin-poc/src/service"
)

type UserController interface {
	FindAll() []entity.User
	Save(ctx *gin.Context) error
}

type controller struct {
	service service.UserService
}

func New(service service.UserService) UserController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.User {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) error {
	var user entity.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		return err
	}
	c.service.Save(user)
	return nil
}
