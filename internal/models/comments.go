package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Comment struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id" gorm:"size:191"`
	RecipeID  uuid.UUID `json:"recipe_id" gorm:"size:191"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      User      `json:"commenter"`
}

func (c *Comment) BeforeCreate(tx *gorm.DB) error {
	c.ID = uuid.New()
	return nil
}

type CommentResponse struct {
	ID        uuid.UUID `json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      Author    `json:"commenter"`
}

func (c *Comment) MapResponse() CommentResponse {
	author := Author{
		ID:       c.User.ID,
		Username: c.User.Username,
	}

	return CommentResponse{
		ID:        c.ID,
		Content:   c.Content,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
		User:      author,
	}
}
