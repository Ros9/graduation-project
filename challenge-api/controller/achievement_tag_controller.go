package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"graduation-project/challenge-api/model"
	"graduation-project/challenge-api/service"
	"io/ioutil"
)

type AchievementTagController interface {
	CreateAchievementTag() gin.HandlerFunc
	GetTagsIdsByAchievementId() gin.HandlerFunc
}

type achievementTagController struct {
	achievementTagService service.AchievementTagService
}

func NewAchievementTagController(achievementTagService service.AchievementTagService) AchievementTagController {
	return &achievementTagController{achievementTagService}
}

func (ac *achievementTagController) CreateAchievementTag() gin.HandlerFunc {
	return func(context *gin.Context) {
		jsonData, err := ioutil.ReadAll(context.Request.Body)
		if err != nil {
			context.JSON(500, err.Error())
		}
		achievementTag := &model.AchievementTag{}
		err = json.Unmarshal(jsonData, achievementTag)
		if err != nil {
			context.JSON(404, err.Error())
		}
		createdAchievementTag, err := ac.achievementTagService.CreateAchievementTag(achievementTag)
		if err != nil {
			context.JSON(404, err.Error())
		}
		context.JSON(200, createdAchievementTag)
	}
}

func (ac *achievementTagController) GetTagsIdsByAchievementId() gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Param("id")
		ats, err := ac.achievementTagService.GetTagsIdsByAchievementId(id)
		if err != nil {
			context.JSON(404, err.Error())
		}
		context.JSON(200, ats)
	}
}
