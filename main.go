package main

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/issues", getIssues)
	r.Use(static.Serve("/", static.LocalFile("storage", true)))
	r.Run(":8888") // listen and serve on 0.0.0.0:8888
}

func getIssues(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{})
}
