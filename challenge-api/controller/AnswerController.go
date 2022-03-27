package controller

import (
	"github.com/gin-gonic/gin"
	"graduation-project/challenge-api/service"
)

type AnswerController interface {
	CreateAnswer() gin.HandlerFunc
	GetAnswer() gin.HandlerFunc
	GetAnswers() gin.HandlerFunc
	UpdateAnswer() gin.HandlerFunc
	DeleteAnswer() gin.HandlerFunc
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

func (cc *answerController) GetAnswer() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

func (cc *answerController) GetAnswers() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

func (cc *answerController) UpdateAnswer() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

func (cc *answerController) DeleteAnswer() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}
