package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"graduation-project/challenge-api/model"
	"io/ioutil"
	"net/http"
)

type Session struct {
	Username string
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/assets", "./assets")
	router.Static("/assets2", "./assets2")

	httpClt := http.DefaultClient

	sessions := map[string]Session{}

	router.GET("/login", func(context *gin.Context) {
		context.HTML(http.StatusOK, "login.html", gin.H{})
	})

	router.POST("/login", func(context *gin.Context) {
		name := context.PostForm("your_name")
		password := context.PostForm("your_pass")
		userAuthInfo := &model.AuthInfo{
			Login:    name,
			Password: password,
		}
		data, err := json.Marshal(userAuthInfo)
		if err != nil {
			fmt.Println("error =", err)
			return
		}
		fmt.Println(string(data))
		url := "http://localhost:8080/auth/user"
		resp, err := httpClt.Post(url, "application/json", bytes.NewBuffer(data))
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("error =", err.Error())
			return
		}
		var token string
		err = json.Unmarshal(body, &token)
		if err != nil {
			fmt.Print("error =", err.Error())
			return
		}
		fmt.Println("token =", token)
		context.SetCookie("session_token", token, 180, "/", "", true, true)
		sessions[token] = Session{Username: name}
		context.Redirect(301, "/index")
	})

	router.GET("/registration", func(context *gin.Context) {
		context.HTML(200, "registration.html", gin.H{})
	})

	router.POST("/registration", func(context *gin.Context) {
		username := context.PostForm("username")
		password := context.PostForm("password")
		email := context.PostForm("email")
		telegram := context.PostForm("telegram")
		fmt.Println(username)
		fmt.Println(password)
		fmt.Println(email)
		fmt.Println(telegram)
		user := &model.User{
			Username: username,
			Password: password,
			Email:    email,
			Telegram: telegram,
		}
		fmt.Println(*user)
		data, err := json.Marshal(user)
		if err != nil {
			fmt.Println("error =", err.Error())
			return
		}
		fmt.Println(string(data))
		url := "http://localhost:8080/user/registration"
		resp, err := httpClt.Post(url, "application/json", bytes.NewBuffer(data))
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("error =", err.Error())
			return
		}
		fmt.Println("body =", string(body))
	})

	router.GET("/challenge/:id", func(context *gin.Context) {
		challengeId := context.Param("id")
		fmt.Println("id =", challengeId)
		url := "http://localhost:8080/challenge/" + challengeId
		fmt.Println("url =", url)
		resp, err := httpClt.Get(url)
		if err != nil {
			fmt.Println("error =", err.Error())
			return
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("error =", err.Error())
			return
		}
		fmt.Println("body =", string(body))
		challenge := model.Challenge{}
		err = json.Unmarshal(body, &challenge)
		context.HTML(200, "challenge.html", gin.H{
			"challenge": challenge,
		})
	})

	router.POST("/challenge/:id", func(context *gin.Context) {
		token, err := context.Cookie("session_token")
		if err != nil {
			fmt.Println(err.Error())
		}
		challengeId := context.Param("id")
		answerCode := context.PostForm("answer")
		fmt.Println("token =", token)
		fmt.Println("challengeId =", challengeId)
		fmt.Println("answer =", answerCode)
		answer := &model.Answer{
			ChallengeID: challengeId,
			Answer:      answerCode,
		}
		fmt.Println(*answer)
		data, err := json.Marshal(answer)
		if err != nil {
			fmt.Println("error =", err.Error())
			return
		}
		fmt.Println(string(data))
		url := "http://localhost:8080/answer"
		fmt.Println("url =", url)
		req, _ := http.NewRequest("POST", url, bytes.NewBuffer(data))
		req.Header.Set("Authorization", "Bearer "+token)
		resp, err := httpClt.Do(req)
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("error =", err.Error())
			return
		}
		fmt.Println("body =", string(body))
	})

	router.GET("/index", func(context *gin.Context) {
		url := "http://localhost:8080/challenges"
		resp, err := httpClt.Get(url)
		if err != nil {
			fmt.Println("error =", err.Error())
			return
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("error =", err.Error())
			return
		}
		fmt.Println("body =", string(body))
		challenges := []model.Challenge{}
		err = json.Unmarshal(body, &challenges)
		if err != nil {
			fmt.Println("error =", err.Error())
			return
		}
		context.HTML(http.StatusOK, "index.html", gin.H{
			"challenges": challenges,
		})
	})

	router.GET("/user/challenges", func(context *gin.Context) {
		token, err := context.Cookie("session_token")
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println("token =", token)
		url := "http://localhost:8080/user/info"
		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Set("Authorization", "Bearer "+token)
		fmt.Println("Authorization header =", req.Header.Get("Authorization"))
		resp, err := httpClt.Do(req)
		if err != nil {
			fmt.Println("error =", err.Error())
			return
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("error =", err.Error())
			return
		}
		fmt.Println("body =", string(body))
		userInfo := model.User{}
		err = json.Unmarshal(body, &userInfo)
		if err != nil {
			fmt.Println("error =", err.Error())
			return
		}
		context.HTML(http.StatusOK, "index.html", gin.H{
			"user_info": userInfo,
		})
	})

	router.GET("/welcome", func(context *gin.Context) {
		token, err := context.Cookie("session_token")
		if err != nil {
			fmt.Println(err.Error())
		}
		name, ok := sessions[token]
		if !ok {
			fmt.Println("name has not exist")
		}
		context.JSON(http.StatusOK, gin.H{
			"hello": name,
		})
	})

	router.GET("/test", func(context *gin.Context) {

	})

	router.Run(":8081")
}
