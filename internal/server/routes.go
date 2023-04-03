package server

import (
	"github.com/gin-gonic/gin"
	authHttpHandler "github.com/mrzalr/cookshare-go/internal/auth/delivery/http"
	authUcase "github.com/mrzalr/cookshare-go/internal/auth/usecase"
	recipeHttpHandler "github.com/mrzalr/cookshare-go/internal/recipe/delivery/http"
	recipeRepo "github.com/mrzalr/cookshare-go/internal/recipe/repository/mysql"
	recipeUcase "github.com/mrzalr/cookshare-go/internal/recipe/usecase"
	userHttpHandler "github.com/mrzalr/cookshare-go/internal/user/delivery/http"
	userRepo "github.com/mrzalr/cookshare-go/internal/user/repository/mysql"
	userUcase "github.com/mrzalr/cookshare-go/internal/user/usecase"
)

func (s *server) MapRoutes(app *gin.Engine) {
	// REPOSITORY
	userRepository := userRepo.New(s.DB)
	recipeRepository := recipeRepo.New(s.DB)

	// USECASE
	authUsecase := authUcase.New(userRepository)
	userUsecase := userUcase.New(userRepository)
	recipeUsecase := recipeUcase.New(recipeRepository)

	// HANDLER
	authHandler := authHttpHandler.New(authUsecase)
	userHandler := userHttpHandler.New(userUsecase, recipeUsecase, s.Mw)
	recipeHandler := recipeHttpHandler.New(recipeUsecase, s.Mw)

	// GROUPING
	// VERSIONING
	v1 := app.Group("api/v1")

	// DOMAIN
	authRoutes := v1.Group("/auth")
	userRoutes := v1.Group("/users")
	recipeRoutes := v1.Group("/recipes")

	authHttpHandler.MapHandlers(*authRoutes, *authHandler)
	userHttpHandler.MapHandlers(*userRoutes, *userHandler)
	recipeHttpHandler.MapHandlers(*recipeRoutes, *recipeHandler)
}
