package unodatabase

import "github.com/LuizGuilherme13/unodatabase/pkg/unodatabase/database"

func (db *DB) Model(model any) database.Repository {
	return db.Repository.Model(model)
}
