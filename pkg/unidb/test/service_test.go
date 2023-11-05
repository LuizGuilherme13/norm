package test

import (
	"testing"

	"github.com/LuizGuilherme13/unidb/internal/config"
	"github.com/LuizGuilherme13/unidb/pkg/unidb"
	"github.com/LuizGuilherme13/unidb/pkg/unidb/entity"
	"github.com/LuizGuilherme13/unidb/pkg/unidb/repository/postgres"
)

func TestFindAll(t *testing.T) {

	env, err := config.GetEnv()
	if err != nil {
		t.Error(err)
	}

	conn := entity.Conn{
		Driver:   env["DB_DRIVER"],
		Host:     env["DB_HOST"],
		Port:     env["DB_PORT"],
		User:     env["DB_USER"],
		Password: env["DB_PASSWORD"],
		DBName:   env["DB_NAME"],
		SSLMode:  env["DB_SSLMODE"],
	}
	repo := postgres.NewPostgresRepo(conn)

	uniService := unidb.NewService(repo)

	result, err := uniService.FindAll("SELECT id, name, phone FROM cliente")
	if err != nil {
		t.Error(err)
	}

	t.Log("Success: ", result)
}
