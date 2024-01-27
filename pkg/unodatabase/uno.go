package unodatabase

import (
	"github.com/LuizGuilherme13/unodatabase/pkg/unodatabase/database"
)

// DB is the struct responsible for containing the repository
type DB struct {
	Repository database.Repository
}

// NewDB is a function that creates a new repository instance
func NewDB(repo database.Repository) database.Repository {
	return &DB{
		Repository: repo,
	}
}
