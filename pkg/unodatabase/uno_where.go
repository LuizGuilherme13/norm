package unodatabase

import "github.com/LuizGuilherme13/unodatabase/pkg/unodatabase/database"

func (db *DB) Where(condition string, values ...any) database.Repository {
	return db.Repository.Where(condition, values...)
}
