package http

import (
	"github.com/gin-gonic/gin"
)

func MapHandlers(r gin.RouterGroup, h handler) {
	r.POST("/", h.mw.Auth(), h.CreateRecipe())
	r.GET("/", h.GetAllRecipes())
	r.GET("/:id", h.GetRecipeByID())
	r.PATCH("/:id", h.mw.Auth(), h.UpdateRecipe())
	r.DELETE("/:id", h.mw.Auth(), h.DeleteRecipe())
}
