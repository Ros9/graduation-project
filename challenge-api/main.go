package main

import (
	"github.com/gin-gonic/gin"
	"graduation-project/challenge-api/controller"
	"graduation-project/challenge-api/service"
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
	answerController := controller.NewAnswerController(answerService)
	commentController := controller.NewCommentController(commentService)
	attachmentController := controller.NewAttachmentController(attachmentService)
	tagController := controller.NewTagController(tagService)
	achievementController := controller.NewAchievementController(achievementService)
	bonusController := controller.NewBonusController(bonusService)

	router.Handle("POST", "/user/register", userController.CreateUser())
	router.Handle("GET", "/user/:id", userController.GetUser())
	router.Handle("GET", "/user", userController.GetUserList())
	router.Handle("PUT", "/user/:id", userController.UpdateUser())
	router.Handle("DELETE", "/user/:id", userController.DeleteUser())

	router.Handle("POST", "/company", companyController.CreateCompany())
	router.Handle("GET", "/company/:id", companyController.GetCompany())
	router.Handle("GET", "/company", companyController.GetCompanies())
	router.Handle("PUT", "/company/:id", companyController.UpdateCompany())
	router.Handle("DELETE", "/company/:id", companyController.DeleteCompany())

	router.Handle("POST", "/challenge", challengeController.CreateChallenge())
	router.Handle("GET", "/challenge/:id", challengeController.GetChallenge())
	router.Handle("GET", "/challenge", challengeController.GetChallenges())
	router.Handle("PUT", "/challenge/:id", challengeController.UpdateChallenge())
	router.Handle("DELETE", "/challenge/:id", challengeController.DeleteChallenge())

	router.Handle("POST", "/answer", answerController.CreateAnswer())
	router.Handle("GET", "/answer/:id", answerController.GetAnswer())
	router.Handle("GET", "/answer", answerController.GetAnswers())
	router.Handle("PUT", "/answer/:id", answerController.UpdateAnswer())
	router.Handle("DELETE", "/answer/:id", answerController.DeleteAnswer())

	router.Handle("POST", "/comment", commentController.CreateComment())
	router.Handle("GET", "/comment/:id", commentController.GetComment())
	router.Handle("GET", "/comment", commentController.GetComments())
	router.Handle("PUT", "/comment/:id", commentController.UpdateComment())
	router.Handle("DELETE", "/comment/:id", commentController.DeleteComment())

	router.Handle("POST", "/attachment", attachmentController.UploadAttachment())
	router.Handle("GET", "/attachment/:id", attachmentController.GetAttachment())

	router.Handle("POST", "/tag", tagController.CreateTag())
	router.Handle("GET", "/tag/:id", tagController.GetTag())
	router.Handle("GET", "/tag", tagController.GetTags())
	router.Handle("PUT", "/tag/:id", tagController.UpdateTag())
	router.Handle("DELETE", "/tag/:id", tagController.DeleteTag())

	router.Handle("POST", "/achievement", achievementController.CreateAchievement())
	router.Handle("GET", "/achievement/:id", achievementController.GetAchievement())
	router.Handle("GET", "/achievement", achievementController.GetAchievements())
	router.Handle("PUT", "/achievement/:id", achievementController.UpdateAchievement())
	router.Handle("DELETE", "/achievement/:id", achievementController.DeleteAchievement())

	router.Handle("POST", "/bonus", bonusController.CreateBonus())
	router.Handle("GET", "/bonus/:id", bonusController.GetBonus())
	router.Handle("GET", "/bonus", bonusController.GetBonuses())
	router.Handle("PUT", "/bonus/:id", bonusController.UpdateBonus())
	router.Handle("DELETE", "/bonus/:id", bonusController.DeleteBonus())

	router.Run(":8080")
}