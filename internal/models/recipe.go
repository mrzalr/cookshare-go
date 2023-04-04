package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Recipe struct {
	ID           uuid.UUID `json:"id"`
	UserID       uuid.UUID `json:"user_id" gorm:"size:191"`
	Title        string    `json:"title"`
	Portion      int       `json:"portion"`
	CookingTime  int       `json:"cooking_time"`
	Description  string    `json:"description"`
	Ingredients  string    `json:"ingredients"`
	Instructions string    `json:"instructions"`
	Tags         string    `json:"tags"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	User         User      `json:"author"`
}

func (r *Recipe) BeforeCreate(tx *gorm.DB) error {
	r.ID = uuid.New()
	return nil
}

type RecipeResponse struct {
	ID           uuid.UUID         `json:"id"`
	Title        string            `json:"title"`
	Portion      int               `json:"portion"`
	CookingTime  int               `json:"cooking_time"`
	Description  string            `json:"description,omitempty"`
	Ingredients  string            `json:"ingredients,omitempty"`
	Instructions string            `json:"instructions"`
	Tags         string            `json:"tags"`
	CreatedAt    time.Time         `json:"created_at"`
	UpdatedAt    time.Time         `json:"updated_at"`
	User         Author            `json:"author"`
	Comments     []CommentResponse `json:"comments"`
}

func (r *Recipe) MapResponse() RecipeResponse {
	author := Author{
		ID:       r.UserID,
		Username: r.User.Username,
	}

	return RecipeResponse{
		ID:           r.ID,
		Title:        r.Title,
		Portion:      r.Portion,
		CookingTime:  r.CookingTime,
		Description:  r.Description,
		Ingredients:  r.Ingredients,
		Instructions: r.Instructions,
		Tags:         r.Tags,
		CreatedAt:    r.CreatedAt,
		UpdatedAt:    r.UpdatedAt,
		User:         author,
	}
}
