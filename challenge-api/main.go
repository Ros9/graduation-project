package main

import (
	"database/sql"
	"graduation-project/challenge-api/controller"
	"graduation-project/challenge-api/repository"
	"graduation-project/challenge-api/service"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/prometheus/common/log"
)

func main() {
	router := gin.Default()

	dbConnString := "postgres://postgres:1234@localhost:5432/CityGoDB?sslmode=disable" //Alibi :)
	//dbConnString := "postgres://postgres:@localhost:5432/citygodb?sslmode=disable"
	dbConnection, err := sql.Open("postgres", dbConnString)

	if err != nil {
		log.Error(err.Error())
	}
	dbConnection.SetMaxOpenConns(10)

	userRepository := repository.NewUserRepository(dbConnection)
	companyRepository := repository.NewCompanyRepository(dbConnection)
	challengeRepository := repository.NewChallengeRepository(dbConnection)
	answerRepository := repository.NewAnswerRepository(dbConnection)
	commentRepository := repository.NewCommentRepository(dbConnection)
	attachmentRepository := repository.NewAttachmentRepository(dbConnection)
	tagRepository := repository.NewTagRepository(dbConnection)
	achievementRepository := repository.NewAchievementRepository(dbConnection)
	bonusRepository := repository.NewBonusRepository(dbConnection)
	achievementTagRepository := repository.NewAchievementTagRepository(dbConnection)
	challengeTagRepository := repository.NewChallengeTagRepository(dbConnection)
	userTagRepository := repository.NewUserTagRepository(dbConnection)
	userAchievementRepository := repository.NewUserAchievementRepository(dbConnection)

	attachmentService := service.NewAttachmentService(attachmentRepository)
	companyService := service.NewCompanyService(companyRepository, attachmentService)
	userService := service.NewUserService(userRepository, challengeRepository, attachmentService,
		userAchievementRepository, achievementRepository)
	challengeService := service.NewChallengeService(challengeRepository, attachmentService,
		challengeTagRepository, tagRepository)
	answerService := service.NewAnswerService(answerRepository, challengeRepository, challengeTagRepository, achievementRepository,
		userTagRepository, achievementTagRepository, userAchievementRepository)
	commentService := service.NewCommentService(commentRepository)
	tagService := service.NewTagService(tagRepository)
	achievementService := service.NewAchievementService(achievementRepository, achievementTagRepository, tagRepository)
	bonusService := service.NewBonusService(bonusRepository)
	achievementTagService := service.NewAchievementTagService(achievementTagRepository)

	userController := controller.NewUserController(userService)
	companyController := controller.NewCompanyController(companyService)
	challengeController := controller.NewChallengeController(challengeService)
	answerController := controller.NewAnswerController(answerService)
	commentController := controller.NewCommentController(commentService)
	attachmentController := controller.NewAttachmentController(attachmentService)
	tagController := controller.NewTagController(tagService)
	achievementController := controller.NewAchievementController(achievementService)
	bonusController := controller.NewBonusController(bonusService)
	authController := controller.NewAuthController(userService)
	achievementTagController := controller.NewAchievementTagController(achievementTagService)

	router.Handle("POST", "/auth/user", authController.GetUserToken())

	router.Handle("POST", "/user/registration", userController.CreateUser())
	router.Handle("GET", "/user/info", userController.GetUserInfo())
	router.Handle("GET", "/user/:id", userController.GetUser())
	router.Handle("GET", "/user/telegram/:telegram", userController.GetUserByTelegram())
	router.Handle("GET", "/user", userController.GetUserList())
	router.Handle("PUT", "/user/:id", userController.UpdateUser())
	router.Handle("DELETE", "/user/:id", userController.DeleteUser())

	router.Handle("POST", "/company", companyController.CreateCompany())
	router.Handle("GET", "/company/:id", companyController.GetCompany())
	router.Handle("GET", "/companies", companyController.GetCompanies())
	router.Handle("PUT", "/company/:id", companyController.UpdateCompany())
	router.Handle("DELETE", "/company/:id", companyController.DeleteCompany())

	router.Handle("POST", "/challenge", challengeController.CreateChallenge())
	router.Handle("POST", "/challenge/telegram", challengeController.CreateChallengeByTelegram())
	router.Handle("GET", "/challenge/:id", challengeController.GetChallenge())
	router.Handle("GET", "/challenges", challengeController.GetChallenges())
	router.Handle("GET", "/challenges/telegram", challengeController.GetChallengesTgResp())
	router.Handle("GET", "/challenges/user/:userId", challengeController.GetChallengesByUserId())
	router.Handle("GET", "/challenges/telegram/user/:userId", challengeController.GetChallengesTgRespByUserId())
	router.Handle("PUT", "/challenge/:id", challengeController.UpdateChallenge())
	router.Handle("DELETE", "/challenge/:id", challengeController.DeleteChallenge())

	router.Handle("POST", "/answer", answerController.CreateAnswer())
	router.Handle("POST", "/answer/telegram", answerController.PostAnswerFromTelegram())
	//router.Handle("GET", "/answer/:id", answerController.GetAnswer())
	//router.Handle("GET", "/answer", answerController.GetAnswers())
	//router.Handle("PUT", "/answer/:id", answerController.UpdateAnswer())
	//router.Handle("DELETE", "/answer/:id", answerController.DeleteAnswer())

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

	router.POST("/achievement-tag", achievementTagController.CreateAchievementTag())
	router.GET("/tags-by-achievement", achievementTagController.GetTagsIdsByAchievementId())

	router.GET("/health-check", func(context *gin.Context) {
		context.JSON(200, "I am ok!")
	})

	router.Run(":8080")
}
