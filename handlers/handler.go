package handlers

import "github.com/mufti1/simple-rest/models"

// Handler holds the API endpoint's function handler.
type Handler struct {
	mongodb MongoDB
}

// NewHandler function to make connection database into handler
func NewHandler(mongodb MongoDB) *Handler {
	return &Handler{
		mongodb: mongodb,
	}
}

// MongoDB presents the interface for mongodb instance.
type MongoDB interface {
	GetPokedex() ([]models.Pokedex, error)
	InsertPokedex(name string, height float64, ability []string) (*models.Pokedex, error)
}
