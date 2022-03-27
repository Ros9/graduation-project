package main

import (
	"github.com/gin-gonic/gin"
	"graduation-project/challange-api/controller"
)

func main() {
	router := gin.Default()
	userController := controller.NewUserController()
	companyController := controller.NewCompanyController()
	userAnswerController := controller.NewUserAnswerController()
	userController.CreateUser()
	companyController.CreateCompany()
	userAnswerController.CreateUserAnswer()
	router.Handle("GET", "/user", userController.CreateUser())
	router.Run(":8080")
}
