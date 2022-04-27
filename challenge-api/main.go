package main

import (
	"database/sql"
	"fmt"
	"graduation-project/challenge-api/controller"
	"graduation-project/challenge-api/repository"
	"graduation-project/challenge-api/service"
	"graduation-project/challenge-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/prometheus/common/log"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/assets", "./assets")

	dbConnString := "postgres://postgres:1234@localhost:5432/CityGoDB?sslmode=disable"
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
	usersChallengesRepository := repository.NewUsersChallengesRepository(dbConnection)

	userService := service.NewUserService(userRepository, usersChallengesRepository, challengeRepository)
	companyService := service.NewCompanyService(companyRepository)
	challengeService := service.NewChallengeService(challengeRepository)
	answerService := service.NewAnswerService(answerRepository, usersChallengesRepository)
	commentService := service.NewCommentService(commentRepository)
	attachmentService := service.NewAttachmentService(attachmentRepository)
	tagService := service.NewTagService(tagRepository)
	achievementService := service.NewAchievementService(achievementRepository)
	bonusService := service.NewBonusService(bonusRepository)

	userController := controller.NewUserController(userService)
	companyController := controller.NewCompanyController(companyService)
	challengeController := controller.NewChallengeController(challengeService)
	answerController := controller.NewAnswerController(answerService)
	commentController := controller.NewCommentController(commentService)
	attachmentController := controller.NewAttachmentController(attachmentService)
	tagController := controller.NewTagController(tagService)
	achievementController := controller.NewAchievementController(achievementService)
	bonusController := controller.NewBonusController(bonusService)
	authController := controller.NewAuthController(userService, companyService)

	router.Handle("POST", "/auth/user", authController.GetUserToken())
	router.Handle("POST", "/auth/company", authController.GetUserToken())

	router.Handle("POST", "/user/registration", userController.CreateUser())
	router.Handle("GET", "/user/info", userController.GetUserInfo())
	router.Handle("GET", "/user/:id", userController.GetUser())
	router.Handle("GET", "/user/telegram/:telegram", userController.GetUserByTelegram())
	router.Handle("GET", "/user", userController.GetUserList())
	router.Handle("PUT", "/user/:id", userController.UpdateUser())
	router.Handle("DELETE", "/user/:id", userController.DeleteUser())

	router.Handle("POST", "/company/registration", companyController.CreateCompany())
	router.Handle("GET", "/company/:id", companyController.GetCompany())
	router.Handle("GET", "/company", companyController.GetCompanies())
	router.Handle("PUT", "/company/:id", companyController.UpdateCompany())
	router.Handle("DELETE", "/company/:id", companyController.DeleteCompany())

	router.Handle("POST", "/challenge", challengeController.CreateChallenge())
	router.Handle("GET", "/challenge/:id", challengeController.GetChallenge())
	router.Handle("GET", "/challenges", challengeController.GetChallenges())
	router.Handle("GET", "/challenges/user/:userId", challengeController.GetChallengesByUserId())
	router.Handle("PUT", "/challenge/:id", challengeController.UpdateChallenge())
	router.Handle("DELETE", "/challenge/:id", challengeController.DeleteChallenge())

	router.Handle("POST", "/answer", answerController.CreateAnswer())
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

	router.GET("/index", func(c *gin.Context) {
		type Challenge struct {
			Title       string
			Description string
		}
		challenges := []Challenge{
			{
				Title:       "Almaty1",
				Description: "About trains",
			},
			{
				Title:       "Abay",
				Description: "About kazakh poet",
			},
			{
				Title:       "Something",
				Description: "About something",
			},
		}
		c.HTML(http.StatusOK, "index.html", gin.H{
			"challenges": challenges,
		})
	})

	token, err := utils.GetToken("Mukha")
	if err != nil {
		fmt.Println("err =", err.Error())
	}
	fmt.Println("token =", token)
	fmt.Println()
	fmt.Println()
	name, err := utils.ParseToken(token, []byte("qwerty12345"))
	if err != nil {
		fmt.Println("err =", err.Error())
	}
	fmt.Println("name =", name)

	router.Run(":8080")
}
