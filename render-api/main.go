package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/assets", "./assets")
	router.Static("/assets2", "./assets2")

	httpClt := http.DefaultClient

	//router.Handle("GET", "/index", func(context *gin.Context) {
	//	context.HTML(http.StatusOK, "index.html", "")
	//})

	router.GET("/login", func(context *gin.Context) {
		context.HTML(http.StatusOK, "login.html", gin.H{})
	})

	router.GET("/registration", func(context *gin.Context) {
		context.HTML(http.StatusOK, "registration.html", gin.H{})
	})

	router.GET("/index", func(c *gin.Context) {
		resp, err := httpClt.Get("http://localhost:8080/challenge")
		if err != nil {
			fmt.Println("error =", err.Error())
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("error =", err.Error())
		}
		fmt.Println("body =", string(body))
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

	router.Run(":8081")
}
