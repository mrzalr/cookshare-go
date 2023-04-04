package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mrzalr/cookshare-go/internal/middleware"
	"github.com/mrzalr/cookshare-go/internal/models"
	"github.com/mrzalr/cookshare-go/internal/recipe"
	"github.com/mrzalr/cookshare-go/internal/user"
)

type handler struct {
	usecase       user.Usecase
	recipeUsecase recipe.Usecase
	mw            *middleware.MiddlewareManager
}

func New(usecase user.Usecase, recipeUsecase recipe.Usecase, mw *middleware.MiddlewareManager) *handler {
	return &handler{
		usecase:       usecase,
		recipeUsecase: recipeUsecase,
		mw:            mw,
	}
}

func (h *handler) UpdateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// GET USER ID
		userID, err := uuid.Parse(ctx.Value("id").(string))
		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				models.StatusBadRequest([]string{err.Error()}),
			)

			return
		}

		// BINDING JSON BODY TO MODELS
		userRequest := models.User{}
		if err := ctx.ShouldBindJSON(&userRequest); err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				models.StatusBadRequest([]string{err.Error()}),
			)

			return
		}

		// UPDATING USER
		userRequest.ID = userID
		user, err := h.usecase.UpdateUser(userID, userRequest)
		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadGateway,
				models.StatusBadGateway([]string{err.Error()}),
			)

			return
		}

		ctx.JSON(
			http.StatusOK,
			models.StatusCreated(user),
		)
	}
}

func (h *handler) GetUserRecipes() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// GET USER ID
		userID, err := uuid.Parse(ctx.Param("id"))
		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				models.StatusBadRequest([]string{err.Error()}),
			)

			return
		}

		// FETCH USER RECIPES DATA
		recipes, err := h.recipeUsecase.GetRecipeByUser(userID)
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
