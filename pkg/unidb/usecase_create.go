package unidb

import (
	"strconv"
	"strings"

	"github.com/LuizGuilherme13/unodatabase/pkg/internal/structutils"
)

// Create cria um novo registro em UniDB.Table a partir do parâmetro 'model'.
// Este deve ser um ponteiro de struct que corresponde a um registro na tabela.
func (udb *UniDB) Create() error {
	err := udb.Open()
	if err != nil {
		panic(err)
	}
	defer udb.DB.Close()

	tx, err := udb.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, values := structutils.GetTagsNames("db", udb.Query.M, false)
	udb.Query.A = values

	query := "INSERT INTO " + udb.Query.T
	query += " (" + strings.Join(udb.Query.C, ", ") + ") "
	query += "VALUES ("
	for i := range udb.Query.A {
		pos := "$" + strconv.Itoa(i+1) + ", "
		query += pos
	}
	query = strings.TrimSuffix(query, ", ") + ")"

	_, err = tx.Exec(query, udb.Query.A...)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
