package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mrzalr/cookshare-go/internal/middleware"
	"github.com/mrzalr/cookshare-go/internal/models"
	"github.com/mrzalr/cookshare-go/internal/recipe"
)

type handler struct {
	usecase recipe.Usecase
	mw      *middleware.MiddlewareManager
}

func New(usecase recipe.Usecase, mw *middleware.MiddlewareManager) *handler {
	return &handler{usecase, mw}
}

func (h *handler) CreateRecipe() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// BINDING JSON BODY TO MODELS
		recipeRequest := models.Recipe{}
		if err := ctx.ShouldBindJSON(&recipeRequest); err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				models.StatusBadRequest([]string{err.Error()}),
			)

			return
		}

		// GET USER ID
		userID, err := uuid.Parse(ctx.Value("id").(string))
		fmt.Println(userID)
		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				models.StatusBadRequest([]string{err.Error()}),
			)

			return
		}

		// INSERT USER TO DATABASE
		recipeRequest.UserID = userID
		recipe, err := h.usecase.CreateRecipe(recipeRequest)
		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadGateway,
				models.StatusBadGateway([]string{err.Error()}),
			)

			return
		}

		ctx.JSON(
			http.StatusCreated,
			models.StatusCreated(recipe),
		)
	}
}

func (h *handler) GetAllRecipes() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// FETCHING ALL RECIPES
		recipes, err := h.usecase.GetAllRecipes()
		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadGateway,
				models.StatusBadGateway([]string{err.Error()}),
			)

			return
		}

		ctx.JSON(
			http.StatusOK,
			models.StatusOk(recipes),
		)
	}
}

func (h *handler) GetRecipeByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// GET RECIPE ID
		recipeID, err := uuid.Parse(ctx.Param("recipeID"))
		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				models.StatusBadRequest([]string{err.Error()}),
			)

			return
		}

		// FETCHING ALL RECIPES
		recipe, err := h.usecase.GetRecipeByID(recipeID)
		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadGateway,
				models.StatusBadGateway([]string{err.Error()}),
			)

			return
		}

		ctx.JSON(
			http.StatusOK,
			models.StatusOk(recipe),
		)
	}
}

func (h *handler) UpdateRecipe() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// GET RECIPE ID
		recipeID, err := uuid.Parse(ctx.Param("recipeID"))
		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				models.StatusBadRequest([]string{err.Error()}),
			)

			return
		}

		// BINDING JSON BODY TO MODELS
		recipeRequest := models.Recipe{}
		if err := ctx.ShouldBindJSON(&recipeRequest); err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				models.StatusBadRequest([]string{err.Error()}),
			)

			return
		}

		// GET USER ID
		userID, err := uuid.Parse(ctx.Value("id").(string))
		fmt.Println(userID)
		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				models.StatusBadRequest([]string{err.Error()}),
			)

			return
		}

		// UPDATING RECIPE
		recipe, err := h.usecase.UpdateRecipe(recipeID, userID, recipeRequest)
		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadGateway,
				models.StatusBadGateway([]string{err.Error()}),
			)

			return
		}

		ctx.JSON(
			http.StatusOK,
			models.StatusOk(recipe),
		)
	}
}

func (h *handler) DeleteRecipe() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// GET RECIPE ID
		recipeID, err := uuid.Parse(ctx.Param("recipeID"))
		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				models.StatusBadRequest([]string{err.Error()}),
			)

			return
		}

		// GET USER ID
		userID, err := uuid.Parse(ctx.Value("id").(string))
		fmt.Println(userID)
		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				models.StatusBadRequest([]string{err.Error()}),
			)

			return
		}

		// DELETE RECIPE
		err = h.usecase.DeleteRecipe(recipeID, userID)
		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadGateway,
				models.StatusBadGateway([]string{err.Error()}),
			)

			return
		}

		ctx.JSON(
			http.StatusOK,
			models.StatusOk(
				fmt.Sprintf("Recipe with id %v successfully deleted", recipeID),
			),
		)
	}
}
