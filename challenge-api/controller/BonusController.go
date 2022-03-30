package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"graduation-project/challenge-api/model"
	"graduation-project/challenge-api/service"
	"io/ioutil"
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

func (bc *bonusController) CreateBonus() gin.HandlerFunc {
	return func(context *gin.Context) {
		jsonData, err := ioutil.ReadAll(context.Request.Body)
		if err != nil {
			context.JSON(500, err.Error())
		}
		bonus := &model.Bonus{}
		err = json.Unmarshal(jsonData, bonus)
		if err != nil {
			context.JSON(404, err.Error())
		}
		createdBonus, err := bc.bonusService.CreateBonus(bonus)
		if err != nil {
			context.JSON(404, err.Error())
		}
		context.JSON(200, createdBonus)
	}
}

func (bc *bonusController) GetBonus() gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Param("id")
		bonus, err := bc.bonusService.GetBonus(id)
		if err != nil {
			context.JSON(404, err.Error())
		}
		context.JSON(200, bonus)
	}
}

func (bc *bonusController) GetBonuses() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

func (bc *bonusController) UpdateBonus() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

func (bc *bonusController) DeleteBonus() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}
