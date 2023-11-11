package postgres

import (
	"fmt"

	"github.com/LuizGuilherme13/unidb/pkg/unidb"
	"github.com/LuizGuilherme13/unidb/pkg/unidb/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PsqlRepo struct {
	Conn models.DBConn
	DB   *sqlx.DB
}

func NewPostgresRepo(conn models.DBConn) unidb.Manager {
	return &PsqlRepo{
		Conn: conn,
	}
}

func (psql *PsqlRepo) OpenDB() (*sqlx.DB, error) {

	dbx, err := sqlx.Connect(psql.Conn.Driver, psql.Conn.String())
	if err != nil {
		return nil, err
	}

	err = dbx.Ping()
	if err != nil {
		return nil, err
	}

	return dbx, nil
}

func (psql *PsqlRepo) Find(dest interface{}, query string, args ...interface{}) error {
	dbx, err := psql.OpenDB()
	if err != nil {
		return fmt.Errorf("postgres.Find (OpenDB): %w", err)
	}
	defer dbx.Close()

	if err := dbx.Select(dest, query, args...); err != nil {
		return fmt.Errorf("postgres.Find (Select): %w", err)
	}

	dbx.Close()
	return nil
}

func (psql *PsqlRepo) Create() {
}
