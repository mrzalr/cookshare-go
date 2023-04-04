package http

import "github.com/gin-gonic/gin"

func MapHandlers(r *gin.RouterGroup, h *handler) {
	r.PATCH("/:id", h.UpdateComment())
	r.DELETE("/:id", h.DeleteComment())
}
