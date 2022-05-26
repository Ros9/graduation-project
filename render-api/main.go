package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("assets/templates/*")
	router.Static("/assets", "./assets")

	router.Handle("GET", "/signup", func(context *gin.Context) {
		context.HTML(200, "signup.html", nil)
	})

	router.Handle("GET", "/signup2", func(context *gin.Context) {
		context.HTML(200, "signup2.html", nil)
	})

	router.Handle("GET", "/signin", func(context *gin.Context) {
		context.HTML(200, "signin.html", nil)
	})

	router.Handle("GET", "/index", func(context *gin.Context) {
		context.HTML(200, "index.html", nil)
	})

	router.Handle("GET", "/challenges", func(context *gin.Context) {
		context.HTML(200, "quests.html", nil)
	})

	router.Handle("GET", "/challenges/something", func(context *gin.Context) {
		context.HTML(200, "detail-quest.html", nil)
	})

	router.Handle("GET", "/profile", func(context *gin.Context) {
		context.HTML(200, "your-progression.html", nil)
	})

	router.Handle("GET", "/achievements", func(context *gin.Context) {
		context.HTML(200, "achievements.html", nil)
	})

	router.Run(":5000")
}
