package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"graduation-project/challenge-api/model"
	"graduation-project/challenge-api/service"
	"io/ioutil"
)

type ChallengeController interface {
	CreateChallenge() gin.HandlerFunc
	GetChallenge() gin.HandlerFunc
	GetChallenges() gin.HandlerFunc
	GetChallengesTgResp() gin.HandlerFunc
	GetChallengesByUserId() gin.HandlerFunc
	GetChallengesTgRespByUserId() gin.HandlerFunc
	UpdateChallenge() gin.HandlerFunc
	DeleteChallenge() gin.HandlerFunc
}

type challengeController struct {
	challengeService service.ChallengeService
}

func NewChallengeController(challengeService service.ChallengeService) ChallengeController {
	return &challengeController{challengeService}
}

func (cc *challengeController) CreateChallenge() gin.HandlerFunc {
	return func(context *gin.Context) {
		//token := context.GetHeader("Authorization")
		//if token == "" {
		//	context.JSON(403, "token is empty")
		//	return
		//}
		//tokenParts := strings.Split(token, " ")
		//if len(tokenParts) != 2 {
		//	context.JSON(403, "token is incorrect")
		//	return
		//}
		//companyId, err := utils.ParseToken(tokenParts[1], []byte("qwerty12345"))
		//if err != nil {
		//	context.JSON(403, err.Error())
		//	return
		//}
		jsonData, err := ioutil.ReadAll(context.Request.Body)
		if err != nil {
			context.JSON(500, err.Error())
		}
		challenge := &model.Challenge{}
		err = json.Unmarshal(jsonData, challenge)
		if err != nil {
			context.JSON(404, err.Error())
		}
		//challenge.CompanyID = companyId
		createdChallenge, err := cc.challengeService.CreateChallenge(challenge)
		if err != nil {
			context.JSON(404, err.Error())
		}
		context.JSON(200, createdChallenge)
	}
}

func (cc *challengeController) GetChallenge() gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Param("id")
		challenge, err := cc.challengeService.GetChallenge(id)
		if err != nil {
			context.JSON(404, err.Error())
		}
		context.JSON(200, challenge)
	}
}

func (cc *challengeController) GetChallenges() gin.HandlerFunc {
	return func(context *gin.Context) {
		challenges, err := cc.challengeService.GetChallenges()
		if err != nil {
			context.JSON(404, err.Error())
		}
		context.JSON(200, challenges)
	}
}

func (cc *challengeController) GetChallengesTgResp() gin.HandlerFunc {
	return func(context *gin.Context) {
		challenges, err := cc.challengeService.GetChallengesTgResp()

		for _, x := range challenges {
			fmt.Println(x)
		}

		if err != nil {
			context.JSON(404, err.Error())
		}
		context.JSON(200, challenges)
	}
}

func (cc *challengeController) GetChallengesByUserId() gin.HandlerFunc {
	return func(context *gin.Context) {
		userId := context.Param("userId")
		challenges, err := cc.challengeService.GetChallengesByUserId(userId)
		if err != nil {
			context.JSON(404, err.Error())
		}
		context.JSON(200, challenges)
	}
}

func (cc *challengeController) GetChallengesTgRespByUserId() gin.HandlerFunc {
	return func(context *gin.Context) {
		userId := context.Param("userId")
		challenges, err := cc.challengeService.GetChallengesTgRespByUserId(userId)
		if err != nil {
			context.JSON(404, err.Error())
		}
		context.JSON(200, challenges)
	}
}

func (cc *challengeController) UpdateChallenge() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(200, nil)
	}
}

func (cc *challengeController) DeleteChallenge() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(200, nil)
	}
}
