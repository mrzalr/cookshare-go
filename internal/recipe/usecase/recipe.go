package usecase

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/mrzalr/cookshare-go/internal/models"
	"github.com/mrzalr/cookshare-go/internal/recipe"
)

type usecase struct {
	repository recipe.Repository
}

func New(repository recipe.Repository) recipe.Usecase {
	return &usecase{repository}
}

func (u *usecase) CreateRecipe(recipe models.Recipe) (models.Recipe, error) {
	return u.repository.Create(recipe)
}

func (u *usecase) GetAllRecipes() ([]models.ShortRecipe, error) {
	recipes, err := u.repository.Find()
	if err != nil {
		return []models.ShortRecipe{}, err
	}

	shortRecipes := []models.ShortRecipe{}
	for _, recipe := range recipes {
		shortRecipe := recipe.MapToShortVersion()
		shortRecipes = append(shortRecipes, shortRecipe)
	}

	return shortRecipes, nil

}

func (u *usecase) GetRecipeByID(recipeID uuid.UUID) (models.Recipe, error) {
	return u.repository.FindByID(recipeID)
}

func (u *usecase) GetRecipeByUser(userID uuid.UUID) ([]models.Recipe, error) {
	return u.repository.FindByUser(userID)
}

func (u *usecase) UpdateRecipe(recipeID uuid.UUID, userID uuid.UUID, recipe models.Recipe) (models.Recipe, error) {
	// CHECKING IF RECIPE IS EXIST
	foundRecipe, err := u.repository.FindByID(recipeID)
	if err != nil {
		return models.Recipe{}, err
	}

	if foundRecipe.UserID != userID {
		return models.Recipe{}, fmt.Errorf("Non authorized user")
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
		return models.Recipe{}, err
	}

	return recipe, nil
}

func (u *usecase) DeleteRecipe(recipeID uuid.UUID, userID uuid.UUID) error {
	// CHECKING IF RECIPE IS EXIST
	foundRecipe, err := u.repository.FindByID(recipeID)
	if err != nil {
		return err
	}

	if foundRecipe.UserID == userID {
		return fmt.Errorf("Non authorized user")
	}

	err = u.repository.Delete(recipeID)
	if err != nil {
		return err
	}

	return nil
}
