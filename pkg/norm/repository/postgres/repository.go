package postgres

import (
	"github.com/LuizGuilherme13/norm/pkg/db"
	"github.com/LuizGuilherme13/norm/pkg/norm"
	_ "github.com/lib/pq"
)

type Repository struct {
	Conn db.DBConn
}

func New(conn db.DBConn) norm.Repository {
	return &Repository{
		Conn: conn,
	}
}
