package test

import (
	"fmt"
	"testing"

	"github.com/LuizGuilherme13/unodatabase/internal/db/tables"
	"github.com/LuizGuilherme13/unodatabase/pkg/unodatabase"
	"github.com/LuizGuilherme13/unodatabase/pkg/unodatabase/models"
)

func TestGet(t *testing.T) {
	unidb := unodatabase.New(models.DBConn{
		Driver:   "postgres",
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "postgres",
		DBName:   "unidb",
		SSLMode:  "disable",
	})

	account := tables.Account{}

	unidb.Migrate(&account)

	err := unidb.Model(&account).Where("user_id = ?", 1).Get("*")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("Result =>", account.UserID, account.UserName)
}
