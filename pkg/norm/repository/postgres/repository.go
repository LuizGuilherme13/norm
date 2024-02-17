package postgres

import (
	"github.com/LuizGuilherme13/norm/pkg/norm"
)

type Repository struct {
	Conn norm.Config
}

func New(options ...norm.Option) norm.Repository {

	conn := &norm.Config{
		Host:     "localhost",
		Port:     "5432",
		Username: "postgres",
		Password: "postgres",
		Database: "postgres",
		SSLMode:  "disable",
	}

	for _, opt := range options {
		opt(conn)
	}

	return &Repository{
		Conn: *conn,
	}
}
