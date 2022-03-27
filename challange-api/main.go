package main

import (
	"github.com/gin-gonic/gin"
	"graduation-project/challange-api/controller"
	"graduation-project/challange-api/service"
)

func main() {
	router := gin.Default()

	userService := service.NewUserService()
	companyService := service.NewCompanyService()
	challengeService := service.NewChallengeService()
	answerService := service.NewAnswerService()
	commentService := service.NewCommentService()
	attachmentService := service.NewAttachmentService()
	tagService := service.NewTagService()
	achievementService := service.NewAchievementService()
	bonusService := service.NewBonusService()

	userController := controller.NewUserController(userService)
	companyController := controller.NewCompanyController(companyService)
	challengeController := controller.NewChallengeController(challengeService)
	answerContoller := controller.NewAnswerController(answerService)
	commentController := controller.NewCommentController(commentService)
	attchmentController := controller.NewAttachmentController(attachmentService)
	tagController := controller.NewTagController(tagService)
	achievementController := controller.NewAchievementController(achievementService)
	bonusController := controller.NewBonusController(bonusService)

	router.Handle("POST", "/challenge", challengeController.CreateChallenge())
	router.Handle("GET", "/challenge/:id", challengeController.GetChallenge())
	router.Handle("GET", "/challenge", challengeController.GetChallenges())

	router.Run(":8080")
}
