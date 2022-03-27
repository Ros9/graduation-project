package controller

import (
	"github.com/gin-gonic/gin"
	"graduation-project/challange-api/service"
)

type TagController interface {
	GetTag() gin.HandlerFunc
}

type tagController struct {
	tagService service.TagService
}

func NewTagController(tagService service.TagService) TagController {
	return &tagController{tagService}
}

func (cc *tagController) GetTag() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}
