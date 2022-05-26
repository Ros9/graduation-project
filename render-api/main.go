package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"graduation-project/challenge-api/model"
	"io/ioutil"
	"net/http"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("assets/templates/*")
	router.Static("/assets", "./assets")

	httpClt := http.DefaultClient

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
		url := "http://localhost:8080/challenges"
		req, _ := http.NewRequest("GET", url, nil)
		resp, err := httpClt.Do(req)
		if err != nil {
			fmt.Println("error when do request:", url, " =", err.Error())
			return
		}
		data, err := ioutil.ReadAll(resp.Body)
		fmt.Println("data from", url, " =", string(data))
		if err != nil {
			fmt.Println("error =", err.Error())
			return
		}
		challenges := []*model.Challenge{}
		err = json.Unmarshal(data, &challenges)
		for _, challenge := range challenges {
			fmt.Println("1")
			fmt.Println("challenge =", challenge)
		}
		if err != nil {
			fmt.Println("error =", err.Error())
			return
		}
		fmt.Println("challenges =", challenges)
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
