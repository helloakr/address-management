package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("./index.html")
	r.GET("/", index)
	r.POST("/login", login)
	r.Run(":8000")
}

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
func login(c *gin.Context) {
	name := c.PostForm("name")
	str1 := "This is login pages" + name
	c.String(http.StatusOK, str1)
}
