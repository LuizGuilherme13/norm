package postgres

import (
	"github.com/LuizGuilherme13/unodatabase/pkg/unodatabase/database"
	"github.com/LuizGuilherme13/unodatabase/pkg/unodatabase/models"
)

type PsqlConn struct {
	Conn  models.DBConn
	Query models.Query
}

func NewPsqlConn(conn models.DBConn) database.Repository {
	return &PsqlConn{
		Conn: conn,
	}
}
