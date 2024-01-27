package unodatabase

import "github.com/LuizGuilherme13/unodatabase/pkg/unodatabase/database"

func (db *DB) WithModel(model any) database.Repository {
	return db.Repository.WithModel(model)
}
