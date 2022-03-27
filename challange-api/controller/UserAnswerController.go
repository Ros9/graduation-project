package controller

import "github.com/gin-gonic/gin"

type UserAnswerController interface {
	CreateUserAnswer() gin.HandlerFunc
}

type userAnswerController struct {
}

func NewUserAnswerController() UserAnswerController {
	return &userAnswerController{}
}

func (cc *userAnswerController) CreateUserAnswer() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}
