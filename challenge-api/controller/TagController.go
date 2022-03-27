package controller

import (
	"github.com/gin-gonic/gin"
	"graduation-project/challenge-api/service"
)

type TagController interface {
	CreateTag() gin.HandlerFunc
	GetTag() gin.HandlerFunc
	GetTags() gin.HandlerFunc
	UpdateTag() gin.HandlerFunc
	DeleteTag() gin.HandlerFunc
}

type tagController struct {
	tagService service.TagService
}

func NewTagController(tagService service.TagService) TagController {
	return &tagController{tagService}
}

func (cc *tagController) CreateTag() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

func (cc *tagController) GetTag() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

func (cc *tagController) GetTags() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

func (cc *tagController) UpdateTag() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

func (cc *tagController) DeleteTag() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}
