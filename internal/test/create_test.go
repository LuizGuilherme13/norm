package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/LuizGuilherme13/unodatabase/internal/db/tables"
	"github.com/LuizGuilherme13/unodatabase/pkg/unodatabase"
	"github.com/LuizGuilherme13/unodatabase/pkg/unodatabase/models"

	"github.com/go-faker/faker/v4"
)

func TestCreate(t *testing.T) {
	unidb := unodatabase.New(models.DBConn{
		Driver:   "postgres",
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "postgres",
		DBName:   "unidb",
		SSLMode:  "disable",
	})

	account := tables.Account{
		UserName:  faker.Name(),
		Password:  faker.Password(),
		Email:     faker.Email(),
		CreatedOn: time.Now(),
	}

	unidb.Migrate(&account)

	err := unidb.Model(&account).Create()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("Result =>", account.UserName)
}
