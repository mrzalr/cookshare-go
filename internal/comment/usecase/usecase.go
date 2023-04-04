package usecase

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/mrzalr/cookshare-go/internal/comment"
	"github.com/mrzalr/cookshare-go/internal/models"
)

type usecase struct {
	repository comment.Repository
}

func New(repository comment.Repository) comment.Usecase {
	return &usecase{repository}
}

func (u *usecase) CreateNewComment(comment models.Comment) (models.CommentResponse, error) {
	comment, err := u.repository.Create(comment)
	if err != nil {
		return models.CommentResponse{}, err
	}

	comment, err = u.repository.FindByID(comment.ID)
	if err != nil {
		return models.CommentResponse{}, err
	}

	return comment.MapResponse(), nil
}

func (u *usecase) UpdateComment(commentID uuid.UUID, userID uuid.UUID, comment models.Comment) (models.CommentResponse, error) {
	foundedComment, err := u.repository.FindByID(commentID)
	if err != nil {
		return models.CommentResponse{}, err
	}

	if foundedComment.UserID != userID {
		return models.CommentResponse{}, fmt.Errorf("non authorized user")
	}

	foundedComment.Content = comment.Content

	updatedComment, err := u.repository.Update(commentID, foundedComment)
	if err != nil {
		return models.CommentResponse{}, err
	}

	return updatedComment.MapResponse(), nil
}

func (u *usecase) DeleteComment(commentID uuid.UUID, userID uuid.UUID) error {
	foundedComment, err := u.repository.FindByID(commentID)
	if err != nil {
		return err
	}

	if foundedComment.UserID != userID {
		return fmt.Errorf("non authorized user")
	}

	return u.repository.Delete(commentID)
}
