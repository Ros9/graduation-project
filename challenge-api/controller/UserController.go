package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"graduation-project/challenge-api/model"
	"graduation-project/challenge-api/service"
	"io/ioutil"
)

type UserController interface {
	CreateUser() gin.HandlerFunc
	GetUser() gin.HandlerFunc
	GetUserList() gin.HandlerFunc
	UpdateUser() gin.HandlerFunc
	DeleteUser() gin.HandlerFunc
}

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &userController{userService}
}

func (uc *userController) CreateUser() gin.HandlerFunc {
	return func(context *gin.Context) {
		jsonData, err := ioutil.ReadAll(context.Request.Body)
		if err != nil {
			context.JSON(404, err.Error())
		}
		user := &model.User{}
		err = json.Unmarshal(jsonData, user)
		if err != nil {
			context.JSON(404, err.Error())
		}
		createdUser, err := uc.userService.CreateUser(user)
		if err != nil {
			context.JSON(404, err.Error())
		}
		context.JSON(200, createdUser)
	}
}

func (uc *userController) GetUser() gin.HandlerFunc {
	return func(context *gin.Context) {
		userId := context.Param("id")
		user, err := uc.userService.GetUser(userId)
		if err != nil {
			context.JSON(404, err.Error())
		}
		context.JSON(200, user)
	}
}

func (uc *userController) GetUserList() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

func (uc *userController) UpdateUser() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

func (uc *userController) DeleteUser() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}
