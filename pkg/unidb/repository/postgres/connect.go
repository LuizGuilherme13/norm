package postgres

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func (r *Repository) Connect() error {

	dataSourceName := "host=" + r.Conn.Host
	dataSourceName += " port=" + r.Conn.Port
	dataSourceName += " user=" + r.Conn.User
	dataSourceName += " password=" + r.Conn.Password
	dataSourceName += " dbname=" + r.Conn.DBName
	dataSourceName += " sslmode=" + r.Conn.SSLMode

	db, err := sql.Open(r.Conn.Driver, dataSourceName)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	r.DB = db

	return nil
}
