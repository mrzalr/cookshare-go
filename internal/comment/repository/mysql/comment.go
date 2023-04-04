package mysql

import (
	"github.com/google/uuid"
	"github.com/mrzalr/cookshare-go/internal/comment"
	"github.com/mrzalr/cookshare-go/internal/models"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) comment.Repository {
	return &repository{db}
}

func (r *repository) Create(comment models.Comment) (models.Comment, error) {
	tx := r.db.Create(&comment)
	return comment, tx.Error
}

func (r *repository) FindByID(commentID uuid.UUID) (models.Comment, error) {
	comment := models.Comment{}
	tx := r.db.Where("id = ?", commentID).Preload("User").First(&comment)
	return comment, tx.Error
}

func (r *repository) FindByRecipe(recipeID uuid.UUID) ([]models.Comment, error) {
	comments := []models.Comment{}
	tx := r.db.Where("recipe_id = ?", recipeID).Preload("User").Find(&comments)
	return comments, tx.Error
}

func (r *repository) Update(commentID uuid.UUID, comment models.Comment) (models.Comment, error) {
	tx := r.db.Model(&models.Comment{}).Where("id = ?", commentID).Updates(&comment)
	return comment, tx.Error
}

func (r *repository) Delete(commentID uuid.UUID) error {
	tx := r.db.Where("id = ?", commentID).Delete(&models.Comment{})
	return tx.Error
}
