package controller

import (
	"github.com/gin-gonic/gin"
	"graduation-project/challenge-api/service"
)

type AttachmentController interface {
	UploadAttachment() gin.HandlerFunc
	GetAttachment() gin.HandlerFunc
}

type attachmentController struct {
	attachmentService service.AttachmentService
}

func NewAttachmentController(attachmentService service.AttachmentService) AttachmentController {
	return &attachmentController{attachmentService}
}

func (cc *attachmentController) UploadAttachment() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

func (cc *attachmentController) GetAttachment() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}
