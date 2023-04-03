package server

import (
	"github.com/gin-gonic/gin"
	authHttpHandler "github.com/mrzalr/cookshare-go/internal/auth/delivery/http"
	authUsecase "github.com/mrzalr/cookshare-go/internal/auth/usecase"
	authRepo "github.com/mrzalr/cookshare-go/internal/user/repository/mysql"
)

func (s *server) MapHandlers(app *gin.Engine) {
	// REPOSITORY
	authRepository := authRepo.New(s.DB)

	// USECASE
	authUsecase := authUsecase.New(authRepository)

	// HANDLER
	authHandler := authHttpHandler.New(authUsecase)

	// GROUPING
	// VERSIONING
	v1 := app.Group("api/v1")

	// DOMAIN
	authRoutes := v1.Group("/auth")

	authHttpHandler.MapRoutes(*authRoutes, *authHandler)
}
