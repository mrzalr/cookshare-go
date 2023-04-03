package mysql

import (
	"github.com/google/uuid"
	"github.com/mrzalr/cookshare-go/internal/models"
	"github.com/mrzalr/cookshare-go/internal/recipe"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) recipe.Repository {
	return &repository{db}
}

func (r *repository) Create(recipe models.Recipe) (models.Recipe, error) {
	tx := r.db.Create(&recipe)
	return recipe, tx.Error
}

func (r *repository) Find() ([]models.Recipe, error) {
	recipes := []models.Recipe{}
	tx := r.db.Preload("User").Find(&recipes)
	return recipes, tx.Error
}

func (r *repository) FindByID(id uuid.UUID) (models.Recipe, error) {
	recipe := models.Recipe{}
	tx := r.db.Where("id = ?", id).Preload("User").First(&recipe)
	return recipe, tx.Error
}

func (r *repository) FindByUser(userID uuid.UUID) ([]models.Recipe, error) {
	recipe := []models.Recipe{}
	tx := r.db.Where("user_id = ?", userID).Preload("User").First(&recipe)
	return recipe, tx.Error
}

func (r *repository) Update(recipeID uuid.UUID, recipe models.Recipe) (models.Recipe, error) {
	tx := r.db.Where("id = ?", recipeID).Updates(recipe)
	return recipe, tx.Error
}

func (r *repository) Delete(recipeID uuid.UUID) error {
	tx := r.db.Where("id = ?", recipeID).Delete(&models.Recipe{})
	return tx.Error
}
