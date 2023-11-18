package unidb

import (
	"errors"

	"github.com/LuizGuilherme13/unidb/pkg/structutils"
	"github.com/LuizGuilherme13/unidb/pkg/unidb/models"
)

func (udb *UniDB) Get(dest interface{}, query models.QueryBuilder) error {

	err := udb.Open()
	if err != nil {
		panic(err)
	}
	defer udb.DB.Close()

	rows, err := udb.DB.Query(query.Q, query.A...)
	if err != nil {
		return err
	}
	defer rows.Close()

	if rows.Next() {
		err := structutils.Scan(rows, dest)
		if err != nil {
			return err
		}
	} else {
		return errors.New("register not found")
	}

	return nil
}
