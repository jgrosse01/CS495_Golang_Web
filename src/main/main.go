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
			"tab_title":      "TEST PAGE",
			"title":          "EET'S WOOURKING!!!!",
			"footer_content": "FOOTER",
		})
	})

	router.GET("/ping", func(c *gin.Context) {
		c.HTML(http.StatusOK, "pong.tmpl", gin.H{
			"tab_title": "Ping",
			"title":     "Current Time",
			"content":   time.Now(),
		})
	})

	err := router.Run()
	if err != nil {
		fmt.Println("[ERROR] Could not run web application.")
	}
}
