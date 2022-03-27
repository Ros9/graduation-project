package controller

import "github.com/gin-gonic/gin"

type UserController interface {
	CreateUser() gin.HandlerFunc
}

type userController struct {
}

func NewUserController() UserController {
	return &userController{}
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
