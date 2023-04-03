package http

import "github.com/mrzalr/cookshare-go/internal/recipe"

type handler struct {
	usecase recipe.Usecase
}

func New(usecase recipe.Usecase) *handler {
	return &handler{usecase}
}
