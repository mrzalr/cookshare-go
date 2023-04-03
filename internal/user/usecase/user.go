package User

import (
	"github.com/google/uuid"
	"github.com/mrzalr/cookshare-go/internal/models"
	"github.com/mrzalr/cookshare-go/internal/user"
)

type usecase struct {
	repository user.Repository
}

func New(repository user.Repository) user.Usecase {
	return &usecase{repository}
}

func (u *usecase) UpdateUser(userID uuid.UUID, user models.User) (models.User, error) {
	// CHECKING IF USER IS EXIST
	foundUser, err := u.repository.FindByID(userID)
	if err != nil {
		return models.User{}, err
	}

	foundUser.Username = user.Username
	foundUser.Email = user.Email
	foundUser.Password = user.Password

	err = foundUser.HashPassword()
	if err != nil {
		return models.User{}, err
	}

	foundUser.Sanitize()

	user, err = u.repository.Update(userID, foundUser)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
