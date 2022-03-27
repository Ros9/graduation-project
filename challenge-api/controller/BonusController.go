package controller

import (
	"github.com/gin-gonic/gin"
	"graduation-project/challenge-api/service"
)

type BonusController interface {
	CreateBonus() gin.HandlerFunc
	GetBonus() gin.HandlerFunc
	GetBonuses() gin.HandlerFunc
	UpdateBonus() gin.HandlerFunc
	DeleteBonus() gin.HandlerFunc
}

type bonusController struct {
	bonusService service.BonusService
}

func NewBonusController(bonusService service.BonusService) BonusController {
	return &bonusController{bonusService}
}

func (cc *bonusController) CreateBonus() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

func (cc *bonusController) GetBonus() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

func (cc *bonusController) GetBonuses() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

func (cc *bonusController) UpdateBonus() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

func (cc *bonusController) DeleteBonus() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}
