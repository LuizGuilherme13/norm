package main

import (
	"fmt"

	"github.com/LuizGuilherme13/unodatabase/internal/entity"
	"github.com/LuizGuilherme13/unodatabase/pkg/unodatabase"
	"github.com/LuizGuilherme13/unodatabase/pkg/unodatabase/database/postgres"
	"github.com/LuizGuilherme13/unodatabase/pkg/unodatabase/models"
)

func main() {
	uno := unodatabase.NewDB(postgres.NewPsqlConn(models.DBConn{
		Driver:   "postgres",
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "postgres",
		DBName:   "unidb",
		SSLMode:  "disable",
	}))

	a := entity.Account{}

	uno.Migrate(&a)

	err := uno.Model(&a).Where("user_id = $1", 1).Get("*")
	if err != nil {
		panic(err)
	}

	fmt.Println(a)

}

//* OK unidb.Migrate(&account)

// TODO unidb.Model(&account).Where("user_id = ?", 1).Update("username")

// TODO unidb.Model(&account).Where("user_id = ?", 1).Delete()

//* OK unidb.Model(&account).Create()

//* OK unidb.Model(&account).Where("user_id = ?", 1).Get("*")
