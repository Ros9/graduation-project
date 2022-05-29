package controller

import (
	"encoding/json"
	"fmt"
	"graduation-project/challenge-api/model"
	"graduation-project/challenge-api/service"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AttachmentController interface {
	UploadAttachment() gin.HandlerFunc
	UploadAttachmentFromTelegram() gin.HandlerFunc
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

func (ac *attachmentController) UploadAttachmentFromTelegram() gin.HandlerFunc {
	return func(context *gin.Context) {
		jsonData, err := ioutil.ReadAll(context.Request.Body)
		if err != nil {
			context.JSON(500, err.Error())
		}
		attachmentLinkReq := &model.AttachmentLinkReq{}
		err = json.Unmarshal(jsonData, attachmentLinkReq)
		if err != nil {
			context.JSON(404, err.Error())
		}
		fmt.Println("\n\n\nATTTTTAAACHH ", attachmentLinkReq)
		//=====================================================
		// todo@Ros9 Скачать по ссылке и сунуть в бд
		//=====================================================
		if err != nil {
			context.JSON(404, err.Error())
		}
		//context.JSON(200, createdAttachment)
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
