package postgres

import (
	"github.com/LuizGuilherme13/unodatabase/pkg/internal/structutils"
	"github.com/LuizGuilherme13/unodatabase/pkg/unodatabase/database"
)

func (db *PsqlConn) WithModel(model any) database.Repository {
	db.Query.Dest.Self = model
	db.Query.Dest.Fields = structutils.GetTags("db", model)

	modelValues := structutils.GetValues("db", model)
	values := make([]any, 0, len(db.Query.Dest.Fields))

	for _, v := range db.Query.Dest.Fields {
		values = append(values, modelValues[v])
	}
	db.Query.Dest.Values = values

	return db
}
