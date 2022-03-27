package controller

import "github.com/gin-gonic/gin"

type BonusController interface {
	CreateBonus() gin.HandlerFunc
}

type bonusController struct {
}

func NewBonusController() BonusController {
	return &bonusController{}
}

func (cc *bonusController) CreateBonus() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}
