package controller

import (
	"github.com/gin-gonic/gin"
	"graduation-project/challenge-api/service"
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

	}
}

func (ac *achievementController) GetAchievement() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

func (ac *achievementController) GetAchievements() gin.HandlerFunc {
	return func(context *gin.Context) {

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
