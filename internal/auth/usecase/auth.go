package usecase

import (
	"github.com/mrzalr/cookshare-go/internal/auth"
	"github.com/mrzalr/cookshare-go/internal/models"
	"github.com/mrzalr/cookshare-go/internal/user"
)

type usecase struct {
	repository user.Repository
}

func New(repository user.Repository) auth.Usecase {
	return &usecase{repository}
}

func (u *usecase) Register(user models.User) (models.User, error) {
	user, err := u.repository.Create(user)
	if err != nil {
		return models.User{}, err
	}

	user.Sanitize()
	return user, nil
}

func (u *usecase) Login(user models.User) (models.UserWithToken, error) {
	foundUser, err := u.repository.FindByEmail(user.Email)
	if err != nil {
		return models.UserWithToken{}, err
	}

	err = foundUser.ComparePassword(user.Password)
	if err != nil {
		return models.UserWithToken{}, err
	}

	token, err := foundUser.GenerateToken()
	if err != nil {
		return models.UserWithToken{}, err
	}

	foundUser.Sanitize()

	userWithToken := models.UserWithToken{
		User:  foundUser,
		Token: token,
	}

	return userWithToken, nil
}
