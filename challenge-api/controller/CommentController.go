package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"graduation-project/challenge-api/model"
	"graduation-project/challenge-api/service"
	"io/ioutil"
)

type CommentController interface {
	CreateComment() gin.HandlerFunc
	GetComment() gin.HandlerFunc
	GetComments() gin.HandlerFunc
	UpdateComment() gin.HandlerFunc
	DeleteComment() gin.HandlerFunc
}

type commentController struct {
	commentService service.CommentService
}

func NewCommentController(commentService service.CommentService) CommentController {
	return &commentController{commentService}
}

func (cc *commentController) CreateComment() gin.HandlerFunc {
	return func(context *gin.Context) {
		jsonData, err := ioutil.ReadAll(context.Request.Body)
		if err != nil {
			context.JSON(404, err.Error())
		}
		comment := &model.Comment{}
		err = json.Unmarshal(jsonData, comment)
		if err != nil {
			context.JSON(404, err.Error())
		}
		createdComment, err := cc.commentService.CreateComment(comment)
		if err != nil {
			context.JSON(404, err.Error())
		}
		context.JSON(200, createdComment)
	}
}

func (cc *commentController) GetComment() gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Param("id")
		comment, err := cc.commentService.GetComment(id)
		if err != nil {
			context.JSON(404, err.Error())
		}
		context.JSON(200, comment)
	}
}

func (cc *commentController) GetComments() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

func (cc *commentController) UpdateComment() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

func (cc *commentController) DeleteComment() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}
