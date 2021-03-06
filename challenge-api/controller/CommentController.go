package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"graduation-project/challenge-api/model"
	"graduation-project/challenge-api/service"
	"graduation-project/challenge-api/utils"
	"io/ioutil"
	"strings"
)

type CommentController interface {
	CreateComment() gin.HandlerFunc
	GetComment() gin.HandlerFunc
	GetComments() gin.HandlerFunc
	GetCommentsByChallengeId() gin.HandlerFunc
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
		token := context.GetHeader("Authorization")
		if token == "" {
			context.JSON(403, "token is empty")
			return
		}
		tokenParts := strings.Split(token, " ")
		if len(tokenParts) != 2 {
			context.JSON(403, "token is incorrect")
			return
		}
		userId, err := utils.ParseToken(tokenParts[1], []byte("qwerty12345"))
		if err != nil {
			context.JSON(403, err.Error())
			return
		}
		jsonData, err := ioutil.ReadAll(context.Request.Body)
		if err != nil {
			context.JSON(404, err.Error())
		}
		comment := &model.Comment{}
		err = json.Unmarshal(jsonData, comment)
		if err != nil {
			context.JSON(404, err.Error())
		}
		comment.UserId = userId
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

func (cc *commentController) GetCommentsByChallengeId() gin.HandlerFunc {
	return func(context *gin.Context) {
		challengeId := context.Param("challenge_id")
		comments, err := cc.commentService.GetCommentsByChallengeId(challengeId)
		if err != nil {
			fmt.Println("error =", err.Error())
			context.JSON(500, err.Error())
		}
		context.JSON(200, comments)
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
