package postgres

import (
	"strings"

	"github.com/LuizGuilherme13/unodatabase/pkg/internal/structutils"
)

func (db *PsqlConn) Migrate(model any) {
	db.Query.Table = strings.ToLower(structutils.GetName(model))
}
