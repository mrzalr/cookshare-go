package comment

import (
	"github.com/google/uuid"
	"github.com/mrzalr/cookshare-go/internal/models"
)

type Repository interface {
	Create(comment models.Comment) (models.Comment, error)
	FindByID(commentID uuid.UUID) (models.Comment, error)
	FindByRecipe(recipeID uuid.UUID) ([]models.Comment, error)
	Update(commentID uuid.UUID, comment models.Comment) (models.Comment, error)
	Delete(commentID uuid.UUID) error
}

type Usecase interface {
	CreateNewComment(comment models.Comment) (models.CommentResponse, error)
	UpdateComment(commentID uuid.UUID, comment models.Comment) (models.Comment, error)
	DeleteComment(commentID uuid.UUID) error
}
