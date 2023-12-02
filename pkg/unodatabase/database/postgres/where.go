package postgres

import (
	"fmt"

	"github.com/LuizGuilherme13/unodatabase/pkg/unodatabase/database"
	"github.com/LuizGuilherme13/unodatabase/pkg/unodatabase/models"
)

func (db *PsqlConn) Where(condition string, values ...any) database.Repository {
	c := fmt.Sprintf("WHERE %s", condition)
	cond := models.Condition{Condition: c}
	cond.Value = append(cond.Value, values...)

	db.Query.Conditions = append(db.Query.Conditions, cond)

	return db
}
