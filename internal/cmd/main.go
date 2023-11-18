package main

import (
	"github.com/LuizGuilherme13/unidb/internal/models"
	"github.com/LuizGuilherme13/unidb/pkg/unidb"
	modelspkg "github.com/LuizGuilherme13/unidb/pkg/unidb/models"
)

func main() {
	unidb := unidb.New(modelspkg.DBConn{
		Driver:   "postgres",
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "postgres",
		DBName:   "unidb",
		SSLMode:  "disable",
	})

	account := models.Account{
		UserName: "teste",
		Password: "abcdefg",
		Email:    "email2@teste.com",
	}

	if err := unidb.Create(&account); err != nil {
		panic(err)
	}

}
