package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mrzalr/cookshare-go/internal/comment"
	"github.com/mrzalr/cookshare-go/internal/middleware"
	"github.com/mrzalr/cookshare-go/internal/models"
)

type handler struct {
	usecase comment.Usecase
	mw      *middleware.MiddlewareManager
}

func New(usecase comment.Usecase, mw *middleware.MiddlewareManager) *handler {
	return &handler{
		usecase: usecase,
		mw:      mw,
	}
}

func (h *handler) UpdateComment() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// GET COMMENT ID
		commentID, err := uuid.Parse(ctx.Param("id"))
		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				models.StatusBadRequest([]string{err.Error()}),
			)

			return
		}

		// BINDING JSON BODY TO MODELS
		commentRequest := models.Comment{}
		if err := ctx.ShouldBindJSON(&commentRequest); err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				models.StatusBadRequest([]string{err.Error()}),
			)

			return
		}

		// GET USER ID
		userID, err := uuid.Parse(ctx.Value("id").(string))
		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				models.StatusBadRequest([]string{err.Error()}),
			)

			return
		}

		// UPDATING RECIPE
		comment, err := h.usecase.UpdateComment(commentID, userID, commentRequest)
		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadGateway,
				models.StatusBadGateway([]string{err.Error()}),
			)

			return
		}

		ctx.JSON(
			http.StatusOK,
			models.StatusOk(comment),
		)
	}
}

func (h *handler) DeleteComment() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// GET COMMENT ID
		commentID, err := uuid.Parse(ctx.Param("id"))
		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				models.StatusBadRequest([]string{err.Error()}),
			)

			return
		}

		// GET USER ID
		userID, err := uuid.Parse(ctx.Value("id").(string))
		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				models.StatusBadRequest([]string{err.Error()}),
			)

			return
		}

		// UPDATING RECIPE
		err = h.usecase.DeleteComment(commentID, userID)
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
				fmt.Sprintf("Comment with id %v successfully deleted", commentID),
			),
		)
	}
}
