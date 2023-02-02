package server

import (
	"net/http"
	handlers "we_bid/api"
	"we_bid/services"

	"github.com/gin-gonic/gin"
)

type server struct {
	router         *gin.Engine
	productHandler handlers.ProductHandler
}

func InitServer(router *gin.Engine) *server {
	productService := services.CreateProductService()
	productHandler := handlers.CreateProductHandler(productService)
	productHandler.Register(router)
	return &server{
		router:         router,
		productHandler: productHandler,
	}
}

func (s *server) Start() error {
	httpServer := &http.Server{
		Handler: s.router,
		Addr:    ":8080",
	}

	return httpServer.ListenAndServe()
}
