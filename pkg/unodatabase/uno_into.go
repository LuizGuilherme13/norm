package unodatabase

import "github.com/LuizGuilherme13/unodatabase/pkg/unodatabase/database"

func (db *DB) Into(table string) database.Repository {
	return db.Repository.Into(table)

}
