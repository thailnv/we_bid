package handlers

import "github.com/gin-gonic/gin"

type Handler interface {
	Register(router gin.IRouter)
}
