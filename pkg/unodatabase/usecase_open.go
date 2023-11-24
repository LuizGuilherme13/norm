package unodatabase

import (
	_ "github.com/lib/pq"

	"database/sql"
)

func (udb *UniDB) Open() error {

	db, err := sql.Open(udb.Conn.Driver, udb.Conn.String())
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	udb.DB = db

	return nil
}
