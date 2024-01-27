package test

import (
	"fmt"
	"testing"

	"github.com/LuizGuilherme13/unodatabase/internal/entity"
	"github.com/LuizGuilherme13/unodatabase/pkg/unodatabase"
	"github.com/LuizGuilherme13/unodatabase/pkg/unodatabase/database/postgres"
	"github.com/LuizGuilherme13/unodatabase/pkg/unodatabase/models"
)

func TestGet(t *testing.T) {
	unoDB := unodatabase.NewDB(postgres.NewPsqlConn(models.DBConn{
		Driver:   "postgres",
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "postgres",
		DBName:   "postgres",
		SSLMode:  "disable",
	}))

	user := entity.User{}

	err := unoDB.Into("users").WithModel(&user).Where("user_id = $1", 1).Find("*")
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("%#v:", user)
}
