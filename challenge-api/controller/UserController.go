package controller

import (
	"encoding/json"
	"fmt"
	"graduation-project/challenge-api/model"
	"graduation-project/challenge-api/service"
	"graduation-project/challenge-api/utils"
	"io/ioutil"
	"strings"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	CreateUser() gin.HandlerFunc
	GetUser() gin.HandlerFunc
	GetUserList() gin.HandlerFunc
	UpdateUser() gin.HandlerFunc
	DeleteUser() gin.HandlerFunc
	GetUserInfo() gin.HandlerFunc
	GetUserByTelegram() gin.HandlerFunc
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

func (uc *userController) GetUserByTelegram() gin.HandlerFunc {
	return func(context *gin.Context) {
		userTelegram := context.Param("telegram")
		fmt.Println("\n\n==== handler", userTelegram)
		user, err := uc.userService.GetUserByTelegram(userTelegram)
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

func (uc *userController) GetUserInfo() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.GetHeader("Authorization")
		if token == "" {
			context.JSON(403, "token is empty")
			return
		}
		tokenParts := strings.Split(token, " ")
		if len(tokenParts) != 2 {
			context.JSON(403, "token is incorrect")
			return
		}
		userId, err := utils.ParseToken(tokenParts[1], []byte("qwerty12345"))
		if err != nil {
			fmt.Println("error when parse token =", err.Error())
			context.JSON(403, err.Error())
			return
		}
		user, err := uc.userService.GetUser(userId)
		if err != nil {
			context.JSON(404, err.Error())
		}
		context.JSON(200, user)
	}
}
