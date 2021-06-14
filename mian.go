package main

import (
	"hummingbird/middlewares"
	"hummingbird/router"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middlewares.Cors())
	r.Use(middlewares.Session("topgoer"))
	r.StaticFS("/static", http.Dir("./studentFile"))
	r = router.ControlRouter(r)
	r.Run(":8000")
}
