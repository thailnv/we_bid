package main

import (
	"we_bid/middlewares"
	"we_bid/server"

	"github.com/gin-gonic/gin"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(middlewares.SetRequestID())
	r.Use(middlewares.SetupLog())
	s := server.InitServer(r)
	s.Start()
}
