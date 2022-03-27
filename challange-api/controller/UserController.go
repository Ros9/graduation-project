package controller

import (
	"github.com/gin-gonic/gin"
	"graduation-project/challange-api/service"
)

type UserController interface {
	CreateUser() gin.HandlerFunc
}

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &userController{userService}
}

func (uc *userController) CreateUser() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

func (uc *userController) GetUser() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

func (uc *userController) GetUserList() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}
