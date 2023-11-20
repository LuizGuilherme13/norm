package unidb

import "github.com/LuizGuilherme13/unidb/pkg/internal/structutils"

// Migrate por enquanto apenas pega o nome da struct e passa para UniDB.Table.
// Sendo necessário utilizar este método sempre que for realizar alguma solicitação ao banco.
func (udb *UniDB) Migrate(model interface{}) error {
	tableName := structutils.GetName(model)

	udb.Query.T = tableName
	return nil
}

func (udb *UniDB) Model(model any) *UniDB {
	columns, _ := structutils.GetTagsNames("db", model, false)

	udb.Query.C = columns
	udb.Query.M = model

	return udb
}
