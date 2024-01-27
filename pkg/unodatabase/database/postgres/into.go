package postgres

import (
	"github.com/LuizGuilherme13/unodatabase/pkg/unodatabase/database"
)

func (db *PsqlConn) Into(table string) database.Repository {
	//strings.ToLower(structutils.GetName(table))
	db.Query.Table = table

	return db
}
