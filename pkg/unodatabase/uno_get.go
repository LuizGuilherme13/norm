package unodatabase

import (
	"database/sql"
	"errors"
)

func (db *DB) Get(cols ...string) error {
	err := db.Repository.Get(cols...)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return errors.New("unodatabase.Get: Register not found")
	case err != nil:
		return err

	}

	return nil
}
