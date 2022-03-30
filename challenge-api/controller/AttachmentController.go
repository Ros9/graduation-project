package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"graduation-project/challenge-api/model"
	"graduation-project/challenge-api/service"
	"io/ioutil"
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
		jsonData, err := ioutil.ReadAll(context.Request.Body)
		if err != nil {
			context.JSON(500, err.Error())
		}
		attachment := &model.Attachment{}
		err = json.Unmarshal(jsonData, attachment)
		if err != nil {
			context.JSON(404, err.Error())
		}
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
