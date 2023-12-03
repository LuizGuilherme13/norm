package unodatabase

import (
	"github.com/LuizGuilherme13/unodatabase/pkg/unodatabase/database"
)

type DB struct {
	Repository database.Repository
}

func NewDB(repo database.Repository) database.Repository {
	return &DB{
		Repository: repo,
	}
}
