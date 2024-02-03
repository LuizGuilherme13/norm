package unodatabase

import (
	"github.com/LuizGuilherme13/unodatabase/pkg/unodatabase/models"
)

type Service struct {
	r      Repository
	Query  models.Query
	Errors models.ErrorGroup
}

// NewService is a function that creates a new repository instance
func NewService(r Repository) UseCase {
	return &Service{
		r: r,
	}
}
