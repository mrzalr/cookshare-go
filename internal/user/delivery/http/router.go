package http

import "github.com/gin-gonic/gin"

func MapHandlers(r gin.RouterGroup, h handler) {
	r.PATCH("/", h.mw.Auth(), h.UpdateUser())
	r.GET("/:id/recipes", h.GetUserRecipes())
}
