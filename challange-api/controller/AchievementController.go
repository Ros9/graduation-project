package controller

import (
	"github.com/gin-gonic/gin"
	"graduation-project/challange-api/service"
)

type AchievementController interface {
	GetAchievement() gin.HandlerFunc
}

type achievementController struct {
	achievementService service.AchievementService
}

func NewAchievementController(achievementService service.AchievementService) AchievementController {
	return &achievementController{achievementService}
}

func (cc *achievementController) GetAchievement() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}
