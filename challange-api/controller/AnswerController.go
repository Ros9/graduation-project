package controller

import (
	"github.com/gin-gonic/gin"
	"graduation-project/challange-api/service"
)

type AnswerController interface {
	CreateAnswer() gin.HandlerFunc
}

type answerController struct {
	answerService service.AnswerService
}

func NewAnswerController(answerService service.AnswerService) AnswerController {
	return &answerController{answerService}
}

func (cc *answerController) CreateAnswer() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}
