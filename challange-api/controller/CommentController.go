package controller

import (
	"github.com/gin-gonic/gin"
	"graduation-project/challange-api/service"
)

type CommentController interface {
	CreateComment() gin.HandlerFunc
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
