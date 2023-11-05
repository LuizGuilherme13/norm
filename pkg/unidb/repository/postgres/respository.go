package postgres

import (
	"database/sql"

	"github.com/LuizGuilherme13/unidb/pkg/unidb"
	"github.com/LuizGuilherme13/unidb/pkg/unidb/entity"
)

type Repository struct {
	Conn entity.Conn
	DB   *sql.DB
}

func NewPostgresRepo(conn entity.Conn) unidb.Repository {
	return &Repository{
		Conn: conn,
	}
}
