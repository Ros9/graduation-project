package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"graduation-project/challange-api/model"
	"graduation-project/challange-api/service"
	"io/ioutil"
)

type ChallengeController interface {
	CreateChallenge() gin.HandlerFunc
	GetChallenge() gin.HandlerFunc
	GetChallenges() gin.HandlerFunc
}

type challengeController struct {
	challengeService service.ChallengeService
}

func NewChallengeController(challengeService service.ChallengeService) ChallengeController {
	return &challengeController{challengeService}
}

func (cc *challengeController) CreateChallenge() gin.HandlerFunc {
	return func(context *gin.Context) {
		jsonData, err := ioutil.ReadAll(context.Request.Body)
		if err != nil {

		}
		challenge := &model.Challenge{}
		err = json.Unmarshal(jsonData, challenge)
		if err != nil {

		}
		createdChallenge, err := cc.challengeService.CreateChallenge(challenge)
		if err != nil {

		}
		context.JSON(200, createdChallenge)
	}
}

func (cc *challengeController) GetChallenge() gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Param("id")
		context.JSON(200, id)
	}
}

func (cc *challengeController) GetChallenges() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(200, nil)
	}
}
