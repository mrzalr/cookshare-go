package http

import "github.com/gin-gonic/gin"

func MapHandlers(r gin.RouterGroup, h handler) {
	r.POST("/register", h.Register())
	r.POST("/login", h.Login())
}
