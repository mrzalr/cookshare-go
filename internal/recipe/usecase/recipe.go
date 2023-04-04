package usecase

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/mrzalr/cookshare-go/internal/comment"
	"github.com/mrzalr/cookshare-go/internal/models"
	"github.com/mrzalr/cookshare-go/internal/recipe"
	"gorm.io/gorm"
)

type usecase struct {
	repository        recipe.Repository
	commentRepository comment.Repository
}

func New(repository recipe.Repository, commentRepository comment.Repository) recipe.Usecase {
	return &usecase{
		repository:        repository,
		commentRepository: commentRepository,
	}
}

func (u *usecase) CreateRecipe(recipe models.Recipe) (models.RecipeResponse, error) {
	createdRecipe, err := u.repository.Create(recipe)
	if err != nil {
		return models.RecipeResponse{}, err
	}

	recipe, err = u.repository.FindByID(createdRecipe.ID)
	if err != nil {
		return models.RecipeResponse{}, err
	}

	return recipe.MapResponse(), nil
}

func (u *usecase) GetAllRecipes() ([]models.RecipeResponse, error) {
	recipes, err := u.repository.Find()
	if err != nil {
		return []models.RecipeResponse{}, err
	}

	recipesResponse := []models.RecipeResponse{}
	for _, recipe := range recipes {
		comments, err := u.commentRepository.FindByRecipe(recipe.ID)
		if err != nil && err != gorm.ErrRecordNotFound {
			return []models.RecipeResponse{}, err
		}

		commentsResponse := []models.CommentResponse{}
		for _, comment := range comments {
			commentsResponse = append(commentsResponse, comment.MapResponse())
		}

		recipeResponse := recipe.MapResponse()
		recipeResponse.Comments = commentsResponse
		recipesResponse = append(recipesResponse, recipeResponse)
	}

	return recipesResponse, nil

}

func (u *usecase) GetRecipeByID(recipeID uuid.UUID) (models.RecipeResponse, error) {
	recipe, err := u.repository.FindByID(recipeID)
	if err != nil {
		return models.RecipeResponse{}, err
	}

	comments, err := u.commentRepository.FindByRecipe(recipe.ID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return models.RecipeResponse{}, err
	}

	commentsResponse := []models.CommentResponse{}
	for _, comment := range comments {
		commentsResponse = append(commentsResponse, comment.MapResponse())
	}

	recipeResponse := recipe.MapResponse()
	recipeResponse.Comments = commentsResponse

	return recipeResponse, nil
}

func (u *usecase) GetRecipeByUser(userID uuid.UUID) ([]models.RecipeResponse, error) {
	recipes, err := u.repository.FindByUser(userID)
	if err != nil {
		return []models.RecipeResponse{}, err
	}

	recipesResponse := []models.RecipeResponse{}
	for _, recipe := range recipes {
		comments, err := u.commentRepository.FindByRecipe(recipe.ID)
		if err != nil && err != gorm.ErrRecordNotFound {
			return []models.RecipeResponse{}, err
		}

		commentsResponse := []models.CommentResponse{}
		for _, comment := range comments {
			commentsResponse = append(commentsResponse, comment.MapResponse())
		}

		recipeResponse := recipe.MapResponse()
		recipeResponse.Comments = commentsResponse
		recipesResponse = append(recipesResponse, recipeResponse)
	}

	return recipesResponse, nil
}

func (u *usecase) UpdateRecipe(recipeID uuid.UUID, userID uuid.UUID, recipe models.Recipe) (models.RecipeResponse, error) {
	// CHECKING IF RECIPE IS EXIST
	foundRecipe, err := u.repository.FindByID(recipeID)
	if err != nil {
		return models.RecipeResponse{}, err
	}

	if foundRecipe.UserID != userID {
		return models.RecipeResponse{}, fmt.Errorf("non authorized user")
	}

	foundRecipe.Title = recipe.Title
	foundRecipe.Portion = recipe.Portion
	foundRecipe.CookingTime = recipe.CookingTime
	foundRecipe.Description = recipe.Description
	foundRecipe.Ingredients = recipe.Ingredients
	foundRecipe.Instructions = recipe.Instructions
	foundRecipe.Tags = recipe.Tags

	recipe, err = u.repository.Update(recipeID, foundRecipe)
	if err != nil {
		return models.RecipeResponse{}, err
	}

	comments, err := u.commentRepository.FindByRecipe(recipe.ID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return models.RecipeResponse{}, err
	}

	commentsResponse := []models.CommentResponse{}
	for _, comment := range comments {
		commentsResponse = append(commentsResponse, comment.MapResponse())
	}

	recipeResponse := recipe.MapResponse()
	recipeResponse.Comments = commentsResponse

	return recipeResponse, nil
}

func (u *usecase) DeleteRecipe(recipeID uuid.UUID, userID uuid.UUID) error {
	// CHECKING IF RECIPE IS EXIST
	foundRecipe, err := u.repository.FindByID(recipeID)
	if err != nil {
		return err
	}

	if foundRecipe.UserID != userID {
		return fmt.Errorf("non authorized user")
	}

	return u.repository.Delete(recipeID)
}
