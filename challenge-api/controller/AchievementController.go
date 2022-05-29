package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"graduation-project/challenge-api/model"
	"graduation-project/challenge-api/service"
	"io/ioutil"
)

type AchievementController interface {
	CreateAchievement() gin.HandlerFunc
	GetAchievement() gin.HandlerFunc
	GetAchievements() gin.HandlerFunc
	UpdateAchievement() gin.HandlerFunc
	DeleteAchievement() gin.HandlerFunc
}

type achievementController struct {
	achievementService service.AchievementService
}

func NewAchievementController(achievementService service.AchievementService) AchievementController {
	return &achievementController{achievementService}
}

func (ac *achievementController) CreateAchievement() gin.HandlerFunc {
	return func(context *gin.Context) {
		jsonData, err := ioutil.ReadAll(context.Request.Body)
		if err != nil {
			context.JSON(500, err.Error())
		}
		achievement := &model.Achievement{}
		err = json.Unmarshal(jsonData, achievement)
		if err != nil {
			context.JSON(404, err.Error())
		}
		createdAchievement, err := ac.achievementService.CreateAchievement(achievement)
		if err != nil {
			context.JSON(404, err.Error())
		}
		context.JSON(200, createdAchievement)
	}
}

func (ac *achievementController) GetAchievement() gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Param("id")
		achievement, err := ac.achievementService.GetAchievement(id)
		if err != nil {
			context.JSON(404, err.Error())
		}
		context.JSON(200, achievement)
	}
}

func (ac *achievementController) GetAchievements() gin.HandlerFunc {
	return func(context *gin.Context) {
		achievements, err := ac.achievementService.GetAchievements()
		if err != nil {
			context.JSON(500, err.Error())
		}
		context.JSON(200, achievements)
	}
}

func (ac *achievementController) UpdateAchievement() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

func (ac *achievementController) DeleteAchievement() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}
