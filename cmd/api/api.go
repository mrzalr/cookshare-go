package api

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/mrzalr/cookshare-go/internal/models"
	"github.com/mrzalr/cookshare-go/internal/server"
	"github.com/mrzalr/cookshare-go/pkg/db/mysql"
)

func StartApplication() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error when load .env files | err : %v", err.Error())
	}

	db, err := mysql.New()
	if err != nil {
		log.Fatalf("Error when create new mysql connection | err : %v", err.Error())
	}
	err = db.AutoMigrate(
		&models.User{},
		&models.Recipe{},
		&models.Comment{},
	)
	if err != nil {
		log.Fatalf("Error when auto migrating table | err : %v", err.Error())
	}

	s := server.New(db)
	if err := s.Run(); err != nil {
		log.Fatalf("Error when run the server | err : %v", err.Error())
	}
}
