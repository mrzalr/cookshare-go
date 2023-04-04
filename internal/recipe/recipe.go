package recipe

import (
	"github.com/google/uuid"
	"github.com/mrzalr/cookshare-go/internal/models"
)

type Repository interface {
	Create(recipe models.Recipe) (models.Recipe, error)
	Find() ([]models.Recipe, error)
	FindByID(id uuid.UUID) (models.Recipe, error)
	FindByUser(userID uuid.UUID) ([]models.Recipe, error)
	Update(recipeID uuid.UUID, recipe models.Recipe) (models.Recipe, error)
	Delete(recipeID uuid.UUID) error
}

type Usecase interface {
	CreateRecipe(recipe models.Recipe) (models.RecipeResponse, error)
	GetAllRecipes() ([]models.RecipeResponse, error)
	GetRecipeByID(recipeID uuid.UUID) (models.RecipeResponse, error)
	GetRecipeByUser(userID uuid.UUID) ([]models.RecipeResponse, error)
	UpdateRecipe(recipeID uuid.UUID, userID uuid.UUID, recipe models.Recipe) (models.RecipeResponse, error)
	DeleteRecipe(recipeID uuid.UUID, userID uuid.UUID) error
}
