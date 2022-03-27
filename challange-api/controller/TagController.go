package controller

import "github.com/gin-gonic/gin"

type TagController interface {
	GetTag() gin.HandlerFunc
}

type tagController struct {
}

func NewTagController() TagController {
	return &tagController{}
}

func (cc *tagController) GetTag() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}
