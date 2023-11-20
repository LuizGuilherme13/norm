package unidb

import "github.com/LuizGuilherme13/unidb/pkg/internal/structutils"

func (udb *UniDB) Model(model any) *UniDB {
	columns, _ := structutils.GetTagsNames("db", model, false)

	udb.Query.C = columns
	udb.Query.M = model

	return udb
}
