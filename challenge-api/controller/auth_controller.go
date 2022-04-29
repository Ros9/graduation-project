package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"graduation-project/challenge-api/model"
	"graduation-project/challenge-api/service"
	"io/ioutil"
)

type AuthController interface {
	GetUserToken() gin.HandlerFunc
	GetCompanyToken() gin.HandlerFunc
}

type authController struct {
	userService    service.UserService
	companyService service.CompanyService
}

func NewAuthController(userService service.UserService, companyService service.CompanyService) AuthController {
	return &authController{userService, companyService}
}

func (ac *authController) GetUserToken() gin.HandlerFunc {
	return func(context *gin.Context) {
		jsonData, err := ioutil.ReadAll(context.Request.Body)
		if err != nil {
			context.JSON(500, err.Error())
		}
		userAuth := &model.AuthInfo{}
		err = json.Unmarshal(jsonData, userAuth)
		if err != nil {
			context.JSON(404, err.Error())
		}
		token, err := ac.userService.GetTokenForUser(userAuth.Login, userAuth.Password)
		if err != nil {
			context.JSON(400, "incorrect login or password")
		}
		fmt.Println(" token ===> ", token)
		context.JSON(200, token)
	}
}

func (ac *authController) GetCompanyToken() gin.HandlerFunc {
	return func(context *gin.Context) {
		jsonData, err := ioutil.ReadAll(context.Request.Body)
		if err != nil {
			context.JSON(500, err.Error())
		}
		companyAuth := &model.AuthInfo{}
		err = json.Unmarshal(jsonData, companyAuth)
		if err != nil {
			context.JSON(404, err.Error())
		}
		token, err := ac.companyService.GetTokenForCompany(companyAuth.Login, companyAuth.Password)
		if err != nil {
			context.JSON(400, "incorrect login or password")
		}
		context.JSON(200, token)
	}
}
