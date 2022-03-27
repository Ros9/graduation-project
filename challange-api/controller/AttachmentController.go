package controller

import "github.com/gin-gonic/gin"

type AttachmentController interface {
	CreateAttachment() gin.HandlerFunc
}

type attachmentController struct {
}

func NewAttachmentController() AttachmentController {
	return &attachmentController{}
}

func (cc *attachmentController) CreateAttachment() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}
