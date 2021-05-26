package main

import (
	"hummingbird/middlewares"
	"hummingbird/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middlewares.Cors())
	r.Use(middlewares.Session("topgoer"))
	r = router.ControlRouter(r)
	r.Run(":8000")
}
