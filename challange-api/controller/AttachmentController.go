package controller

import (
	"github.com/gin-gonic/gin"
	"graduation-project/challange-api/service"
)

type AttachmentController interface {
	CreateAttachment() gin.HandlerFunc
}

type attachmentController struct {
	attachmentService service.AttachmentService
}

func NewAttachmentController(attachmentService service.AttachmentService) AttachmentController {
	return &attachmentController{attachmentService}
}

func (cc *attachmentController) CreateAttachment() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}
