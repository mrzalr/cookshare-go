package auth

import "github.com/mrzalr/cookshare-go/internal/models"

type Usecase interface {
	Register(user models.User) (models.User, error)
	Login(user models.User) (models.UserWithToken, error)
}
