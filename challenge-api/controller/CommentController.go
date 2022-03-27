package controller

import (
	"github.com/gin-gonic/gin"
	"graduation-project/challenge-api/service"
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

	}
}

func (cc *commentController) GetComment() gin.HandlerFunc {
	return func(context *gin.Context) {

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
