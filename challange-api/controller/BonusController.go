package controller

import (
	"github.com/gin-gonic/gin"
	"graduation-project/challange-api/service"
)

type BonusController interface {
	CreateBonus() gin.HandlerFunc
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
