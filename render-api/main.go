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
		fmt.Println("body = ", string(data))
		var token string
		err = json.Unmarshal(body, &token)
		if err != nil {
			fmt.Print("error =", err.Error())
			context.Redirect(301, "/login")
			return
		}
		fmt.Println("token =", token)
		context.SetCookie("session_token", token, 600, "/", "", true, true)
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
		redirectUrl := "/login"
		context.Redirect(301, redirectUrl)
	})

	router.GET("/challenge/:id", func(context *gin.Context) {
		token, err := context.Cookie("session_token")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
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
		userInfo, err := getUserInfoByToken(token)
		context.HTML(200, "challenge.html", gin.H{
			"challenge": challenge,
			"userInfo":  userInfo,
		})
	})

	router.POST("/challenge/:id", func(context *gin.Context) {
		token, err := context.Cookie("session_token")
		if err != nil {
			fmt.Println(err.Error())
			return
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
		redirectUrl := "/challenge/" + challengeId
		context.Redirect(301, redirectUrl)
	})

	router.GET("/index", func(context *gin.Context) {
		token, err := context.Cookie("session_token")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
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
		userInfo, err := getUserInfoByToken(token)
		context.HTML(http.StatusOK, "index.html", gin.H{
			"challenges": challenges,
			"userInfo":   userInfo,
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
		fmt.Println("user =", userInfo)
		context.HTML(http.StatusOK, "user_challenges.html", gin.H{
			"userInfo": userInfo,
		})
	})

	router.GET("/health-check", func(context *gin.Context) {
		context.JSON(200, "I am ok!")
	})

	router.Run(":8081")
}

func getUserInfoByToken(token string) (*model.User, error) {
	url := "http://localhost:8080/user/info"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+token)
	fmt.Println("Authorization header =", req.Header.Get("Authorization"))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("error =", err.Error())
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error =", err.Error())
		return nil, err
	}
	fmt.Println("body =", string(body))
	userInfo := model.User{}
	err = json.Unmarshal(body, &userInfo)
	if err != nil {
		fmt.Println("error =", err.Error())
		return nil, err
	}
	fmt.Println("user =", userInfo)
	return &userInfo, nil
}
