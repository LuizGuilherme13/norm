package main

import (
	"fmt"

	"github.com/LuizGuilherme13/unodatabase/internal/db/tables"
	"github.com/LuizGuilherme13/unodatabase/pkg/unodatabase"
	"github.com/LuizGuilherme13/unodatabase/pkg/unodatabase/models"
)

func main() {
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

	err := unidb.Model(&account).Where("user_id = ?", 6).Get("*")
	if err != nil {
		panic(err)
	}

	fmt.Println(account.UserID, account.UserName)

}

// unidb.Migrate(&account)

// unidb.Model(&account).Where("user_id = ?", 1).Update("username")

// unidb.Model(&account).Create()

// unidb.Model(&account).Where("user_id = ?", 1).Get("*")
