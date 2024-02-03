package postgres

import (
	"github.com/LuizGuilherme13/unodatabase/pkg/unodatabase"
	"github.com/LuizGuilherme13/unodatabase/pkg/unodatabase/models"
	_ "github.com/lib/pq"
)

type Repository struct {
	Conn models.DBConn
}

func New(conn models.DBConn) unodatabase.Repository {
	return &Repository{
		Conn: conn,
	}
}
