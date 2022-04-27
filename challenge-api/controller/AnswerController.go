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

type AnswerController interface {
	CreateAnswer() gin.HandlerFunc
	GetAnswer() gin.HandlerFunc
	GetAnswers() gin.HandlerFunc
	UpdateAnswer() gin.HandlerFunc
	DeleteAnswer() gin.HandlerFunc
	PostAnswerFromTelegram() gin.HandlerFunc
}

type answerController struct {
	answerService service.AnswerService
}

func NewAnswerController(answerService service.AnswerService) AnswerController {
	return &answerController{answerService}
}

func (ac *answerController) CreateAnswer() gin.HandlerFunc {
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
			context.JSON(403, err.Error())
			return
		}
		jsonData, err := ioutil.ReadAll(context.Request.Body)
		if err != nil {
			context.JSON(500, err.Error())
			return
		}
		answer := &model.Answer{}
		err = json.Unmarshal(jsonData, answer)
		if err != nil {
			context.JSON(404, err.Error())
			return
		}
		answer.UserID = userId
		createdAnswer, err := ac.answerService.CreateAnswer(answer)
		if err != nil {
			context.JSON(404, err.Error())
			return
		}
		context.JSON(200, createdAnswer)
	}
}

func (ac *answerController) PostAnswerFromTelegram() gin.HandlerFunc {
	return func(context *gin.Context) {
		jsonData, err := ioutil.ReadAll(context.Request.Body)
		if err != nil {
			context.JSON(500, err.Error())
			return
		}
		answer := model.Answer{}
		err = json.Unmarshal(jsonData, &answer)
		if err != nil {
			fmt.Println("\n\n========= answer", answer)
			context.JSON(404, err.Error())
			return
		}
		result, err := ac.answerService.PostAnswerFromTelegram(&answer)

		if err != nil {
			context.JSON(404, err.Error())
			return
		}
		context.JSON(200, result)

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
