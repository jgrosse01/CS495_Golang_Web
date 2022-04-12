package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("static/templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "page.tmpl", gin.H{
			"tab_title": "Literally A Single Video",
			"title":     "Welcome, Nate",
			"content":   "homepage_content",
		})
	})

	router.GET("/ping", func(c *gin.Context) {
		c.HTML(http.StatusOK, "page.tmpl", gin.H{
			"tab_title": "PONG!",
			"title":     "Current Time",
			"content":   "pong_content",
			"value":     time.Now(),
		})
	})

	err := router.Run()
	if err != nil {
		fmt.Println("[ERROR] Could not run web application.")
	}
}
