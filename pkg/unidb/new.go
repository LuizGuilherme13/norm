package unidb

import (
	"github.com/LuizGuilherme13/unidb/pkg/unidb/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type UniDB struct {
	Conn models.DBConn
	DB   *sqlx.DB
}

func New(conn models.DBConn) *UniDB {
	return &UniDB{
		Conn: conn,
	}
}

func (udb *UniDB) Open() error {

	db, err := sqlx.Connect(udb.Conn.Driver, udb.Conn.String())
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
