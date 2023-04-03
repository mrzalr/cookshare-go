package user

import (
	"github.com/google/uuid"
	"github.com/mrzalr/cookshare-go/internal/models"
)

type Repository interface {
	Create(user models.User) (models.User, error)
	FindByEmail(email string) (models.User, error)
	FindByID(id uuid.UUID) (models.User, error)
	Update(userID uuid.UUID, user models.User) (models.User, error)
}

type Usecase interface {
	UpdateUser(userID uuid.UUID, user models.User) (models.User, error)
}
