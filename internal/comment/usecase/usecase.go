package usecase

import (
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

func (u *usecase) UpdateComment(commentID uuid.UUID, comment models.Comment) (models.Comment, error) {
	foundedComment, err := u.repository.FindByID(commentID)
	if err != nil {
		return models.Comment{}, err
	}

	foundedComment.Content = comment.Content

	return u.repository.Update(commentID, foundedComment)
}

func (u *usecase) DeleteComment(commentID uuid.UUID) error {
	_, err := u.repository.FindByID(commentID)
	if err != nil {
		return err
	}

	return u.repository.Delete(commentID)
}
