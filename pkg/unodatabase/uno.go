package unodatabase

import (
	_ "github.com/lib/pq"

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
