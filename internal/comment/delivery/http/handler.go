package http

import (
	"github.com/gin-gonic/gin"
	"github.com/mrzalr/cookshare-go/internal/comment"
	"github.com/mrzalr/cookshare-go/internal/middleware"
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

	}
}

func (h *handler) DeleteComment() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
