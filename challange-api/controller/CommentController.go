package controller

import "github.com/gin-gonic/gin"

type CommentController interface {
	CreateComment() gin.HandlerFunc
}

type commentController struct {
}

func NewCommentController() CommentController {
	return &commentController{}
}

func (cc *commentController) CreateComment() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}
