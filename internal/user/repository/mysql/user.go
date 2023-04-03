package mysql

import (
	"github.com/google/uuid"
	"github.com/mrzalr/cookshare-go/internal/models"
	"github.com/mrzalr/cookshare-go/internal/user"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.Repository {
	return &repository{db}
}

func (r *repository) Create(user models.User) (models.User, error) {
	tx := r.db.Create(&user)
	return user, tx.Error
}

func (r *repository) FindByEmail(email string) (models.User, error) {
	user := models.User{}
	tx := r.db.Where("email = ?", email).First(&user)
	return user, tx.Error
}

func (r *repository) FindByID(id uuid.UUID) (models.User, error) {
	user := models.User{}
	tx := r.db.Where("id = ?", id).First(&user)
	return user, tx.Error
}

func (r *repository) Update(userID uuid.UUID, user models.User) (models.User, error) {
	tx := r.db.Where("id = ?", userID).Updates(user)
	return user, tx.Error
}
