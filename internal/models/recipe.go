package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Recipe struct {
	ID           uuid.UUID `json:"id"`
	UserID       uuid.UUID `json:"user_id"`
	Title        string    `json:"title"`
	Portion      int       `json:"portion"`
	CookingTime  int       `json:"cooking_time"`
	Description  string    `json:"description"`
	Ingredients  string    `json:"ingredients"`
	Instructions string    `json:"instructions"`
	Tags         string    `json:"tags"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	User         User
}

func (r *Recipe) BeforeCreate(tx *gorm.DB) error {
	r.ID = uuid.New()
	return nil
}

type ShortRecipe struct {
	ID          uuid.UUID `json:"id"`
	UserID      uuid.UUID `json:"user_id"`
	Title       string    `json:"title"`
	Portion     int       `json:"portion"`
	CookingTime int       `json:"cooking_time"`
	Description string    `json:"description"`
	Tags        string    `json:"tags"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	User        User
}

func (r *Recipe) MapToShortVersion() ShortRecipe {
	return ShortRecipe{
		ID:          r.ID,
		UserID:      r.UserID,
		Title:       r.Title,
		Portion:     r.Portion,
		CookingTime: r.CookingTime,
		Description: r.Description,
		Tags:        r.Tags,
		CreatedAt:   r.CreatedAt,
		UpdatedAt:   r.UpdatedAt,
		User:        r.User,
	}
}
