package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"graduation-project/challenge-api/model"
	"graduation-project/challenge-api/service"
	"net/http"
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

func (ac *attachmentController) UploadAttachment() gin.HandlerFunc {
	return func(context *gin.Context) {
		file, _, err := context.Request.FormFile("file")
		if err != nil {
			context.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
			return
		}
		externalId := context.PostForm("external_id")
		fmt.Println("externalId =", externalId)
		attachment := &model.Attachment{}
		attachment.ExternalId = externalId
		attachment.File = &file
		createdAttachment, err := ac.attachmentService.CreateAttachment(attachment)
		if err != nil {
			context.JSON(404, err.Error())
		}
		context.JSON(200, createdAttachment)
	}
}

func (ac *attachmentController) GetAttachment() gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Param("id")
		attachment, err := ac.attachmentService.GetAttachment(id)
		if err != nil {
			context.JSON(404, err.Error())
		}
		context.JSON(200, attachment)
	}
}
