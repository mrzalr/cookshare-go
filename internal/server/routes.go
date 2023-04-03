package server

import (
	"github.com/gin-gonic/gin"
	authHttpHandler "github.com/mrzalr/cookshare-go/internal/auth/delivery/http"
	authUsecase "github.com/mrzalr/cookshare-go/internal/auth/usecase"
	userHttpHandler "github.com/mrzalr/cookshare-go/internal/user/delivery/http"
	userRepo "github.com/mrzalr/cookshare-go/internal/user/repository/mysql"
	userUsecase "github.com/mrzalr/cookshare-go/internal/user/usecase"
)

func (s *server) MapRoutes(app *gin.Engine) {
	// REPOSITORY
	userRepository := userRepo.New(s.DB)

	// USECASE
	authUsecase := authUsecase.New(userRepository)
	userUsecase := userUsecase.New(userRepository)

	// HANDLER
	authHandler := authHttpHandler.New(authUsecase)
	userHandler := userHttpHandler.New(userUsecase, s.Mw)

	// GROUPING
	// VERSIONING
	v1 := app.Group("api/v1")

	// DOMAIN
	authRoutes := v1.Group("/auth")
	userRoutes := v1.Group("/users")

	authHttpHandler.MapHandlers(*authRoutes, *authHandler)
	userHttpHandler.MapHandlers(*userRoutes, *userHandler)
}
