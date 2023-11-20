package unidb

import (
	"errors"
	"strings"

	"github.com/LuizGuilherme13/unidb/pkg/internal/structutils"
)

func (udb *UniDB) Get(cols ...string) error {

	err := udb.Open()
	if err != nil {
		panic(err)
	}
	defer udb.DB.Close()

	query := "SELECT " + strings.Join(cols, ", ")
	query += " FROM " + udb.Query.T
	query += udb.Query.Q

	rows, err := udb.DB.Query(query, udb.Query.A...)
	if err != nil {
		return err
	}
	defer rows.Close()

	if rows.Next() {
		err := structutils.ScanToModel(rows, udb.Query.M)
		if err != nil {
			return err
		}
	} else {
		return errors.New("register not found")
	}

	return nil
}
