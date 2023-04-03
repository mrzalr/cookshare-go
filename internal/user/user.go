package user

import "github.com/mrzalr/cookshare-go/internal/models"

type Repository interface {
	Create(user models.User) (models.User, error)
	FindByEmail(email string) (models.User, error)
}
