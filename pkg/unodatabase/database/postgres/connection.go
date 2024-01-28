package postgres

import (
	_ "github.com/lib/pq"

	"github.com/LuizGuilherme13/unodatabase/pkg/unodatabase/database"
	"github.com/LuizGuilherme13/unodatabase/pkg/unodatabase/models"
)

// PsqlConn is the struct that represents the database repository
type PsqlConn struct {
	Conn  models.DBConn
	Query models.Query
}

// NewPsqlConn creates a new PostgresSQL repository instance
func NewPsqlConn(conn models.DBConn) database.Repository {
	return &PsqlConn{
		Conn: conn,
	}
}
