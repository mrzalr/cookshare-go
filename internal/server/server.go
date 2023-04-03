package server

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type server struct {
	DB  *gorm.DB
	App *gin.Engine
}

func New(db *gorm.DB) *server {
	return &server{
		DB:  db,
		App: gin.Default(),
	}
}

func (s *server) Run() error {
	s.MapHandlers(s.App)

	port := os.Getenv("SERVER_PORT")
	log.Printf("Server running on port %s", port)
	return s.App.Run(port)
}
