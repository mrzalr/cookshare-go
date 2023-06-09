package server

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mrzalr/cookshare-go/internal/middleware"
	"gorm.io/gorm"
)

type server struct {
	DB  *gorm.DB
	App *gin.Engine
	Mw  *middleware.MiddlewareManager
}

func New(db *gorm.DB) *server {
	return &server{
		DB:  db,
		App: gin.Default(),
		Mw:  new(middleware.MiddlewareManager),
	}
}

func (s *server) Run() error {
	s.MapRoutes(s.App)

	port := os.Getenv("SERVER_PORT")
	log.Printf("Server running on port %s", port)
	return s.App.Run(fmt.Sprintf(":%s", port))
}
