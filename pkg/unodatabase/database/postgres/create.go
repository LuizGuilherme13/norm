package postgres

import (
	"database/sql"
	"fmt"
	"strings"
)

func (db *PsqlConn) New() error {
	conn, err := sql.Open(db.Conn.Driver, db.Conn.String())
	if err != nil {
		return err
	}
	defer conn.Close()

	tx, err := conn.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var valuesPosition []string
	for i := range db.Query.Dest.Values {
		// fmt.Println(db.Query.Dest.Values[i])
		fmt.Scanln()
		valuesPosition = append(valuesPosition, fmt.Sprintf("$%d", i+1))
	}

	query := fmt.Sprintf("INSERT INTO %s ", db.Query.Table)
	query += fmt.Sprintf("(%s) ", strings.Join(db.Query.Dest.Fields, ", "))
	query += fmt.Sprintf("VALUES (%s)", strings.Join(valuesPosition, ", "))

	// fmt.Println(query)
	// fmt.Println(values...)

	_, err = tx.Exec(query, db.Query.Dest.Values...)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
