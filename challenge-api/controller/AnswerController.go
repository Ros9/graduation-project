package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"graduation-project/challenge-api/model"
	"graduation-project/challenge-api/service"
	"io/ioutil"
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

func (ac *answerController) CreateAnswer() gin.HandlerFunc {
	return func(context *gin.Context) {
		jsonData, err := ioutil.ReadAll(context.Request.Body)
		if err != nil {
			context.JSON(500, err.Error())
		}
		answer := &model.Answer{}
		err = json.Unmarshal(jsonData, answer)
		if err != nil {
			context.JSON(404, err.Error())
		}
		createdAnswer, err := ac.answerService.CreateAnswer(answer)
		if err != nil {
			context.JSON(404, err.Error())
		}
		context.JSON(200, createdAnswer)
	}
}

func (ac *answerController) GetAnswer() gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Param("id")
		answer, err := ac.answerService.GetAnswer(id)
		if err != nil {
			context.JSON(404, err.Error())
		}
		context.JSON(200, answer)
	}
}

func (ac *answerController) GetAnswers() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

func (ac *answerController) UpdateAnswer() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

func (ac *answerController) DeleteAnswer() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}
