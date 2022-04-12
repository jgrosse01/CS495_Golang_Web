package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// Function to run the Go Web Application
// This is great and all, but take a look at page.tmpl as well, I figured out how tmpl conditionals work :)
func main() {
	// define router and load templates
	router := gin.Default()
	router.LoadHTMLGlob("static/templates/*")

	// web-app paths that provide content to fill into the template
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

	// finally, run the app and check errors
	err := router.Run()
	if err != nil {
		fmt.Println("[ERROR] Could not run web application.")
	}
}
