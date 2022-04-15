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
		fmt.Println("body =", string(body))
		token := string(body)
		context.SetCookie("session_token", token, 180, "/", "", true, true)
		sessions[token] = Session{Username: name}
	})

	router.POST("/registration", func(context *gin.Context) {
		name := context.PostForm("your_name")
		password := context.PostForm("your_pass")
		user := &model.User{
			ID:         "",
			Login:      name,
			Email:      name,
			Name:       name,
			Surname:    name,
			Password:   password,
			Telegram:   name,
			Challenges: nil,
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

	router.GET("/index", func(c *gin.Context) {
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
		c.HTML(http.StatusOK, "index.html", gin.H{
			"challenges": challenges,
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

	router.Run(":8081")
}
