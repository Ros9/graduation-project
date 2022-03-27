package controller

import "github.com/gin-gonic/gin"

type AchievementController interface {
	GetAchievement() gin.HandlerFunc
}

type achievementController struct {
}

func NewAchievementController() AchievementController {
	return &achievementController{}
}

func (cc *achievementController) GetAchievement() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}
