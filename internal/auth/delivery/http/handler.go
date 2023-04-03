package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrzalr/cookshare-go/internal/auth"
	"github.com/mrzalr/cookshare-go/internal/models"
)

type handler struct {
	usecase auth.Usecase
}

func New(usecase auth.Usecase) *handler {
	return &handler{usecase}
}

func (h *handler) Register() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// BINDING JSON BODY TO MODELS
		userRequest := models.User{}
		if err := ctx.ShouldBindJSON(&userRequest); err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				models.StatusBadRequest([]string{err.Error()}),
			)

			return
		}

		// INSERT USER TO DATABASE
		user, err := h.usecase.Register(userRequest)
		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadGateway,
				models.StatusBadGateway([]string{err.Error()}),
			)

			return
		}

		ctx.JSON(
			http.StatusCreated,
			models.StatusCreated(user),
		)

	}
}

func (h *handler) Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// BINDING JSON BODY TO MODELS
		userRequest := models.User{}
		if err := ctx.ShouldBindJSON(&userRequest); err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				models.StatusBadRequest([]string{err.Error()}),
			)

			return
		}

		// CHECK USER CREDENTIALS
		user, err := h.usecase.Login(userRequest)
		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadGateway,
				models.StatusBadGateway([]string{err.Error()}),
			)

			return
		}

		ctx.JSON(
			http.StatusOK,
			models.StatusOk(user),
		)
	}
}
