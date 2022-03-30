package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"graduation-project/challenge-api/model"
	"graduation-project/challenge-api/service"
	"io/ioutil"
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

func (tc *tagController) CreateTag() gin.HandlerFunc {
	return func(context *gin.Context) {
		jsonData, err := ioutil.ReadAll(context.Request.Body)
		if err != nil {
			context.JSON(500, err.Error())
		}
		tag := &model.Tag{}
		err = json.Unmarshal(jsonData, tag)
		if err != nil {
			context.JSON(404, err.Error())
		}
		createdTag, err := tc.tagService.CreateTag(tag)
		if err != nil {
			context.JSON(404, err.Error())
		}
		context.JSON(200, createdTag)
	}
}

func (tc *tagController) GetTag() gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Param("id")
		tag, err := tc.tagService.GetTag(id)
		if err != nil {
			context.JSON(404, err.Error())
		}
		context.JSON(200, tag)
	}
}

func (tc *tagController) GetTags() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

func (tc *tagController) UpdateTag() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

func (tc *tagController) DeleteTag() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}
