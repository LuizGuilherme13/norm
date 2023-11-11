package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/LuizGuilherme13/unidb/pkg/unidb"
	"github.com/LuizGuilherme13/unidb/pkg/unidb/models"
	"github.com/LuizGuilherme13/unidb/pkg/unidb/repository/postgres"
)

type Accounts struct {
	UserID    int64        `db:"user_id"`
	UserName  string       `db:"username"`
	Password  string       `db:"password"`
	Email     string       `db:"email"`
	CreatedOn time.Time    `db:"created_on"`
	LastLogin sql.NullTime `db:"last_login"`
}

func main() {
	unidb := unidb.New(postgres.NewPostgresRepo(models.DBConn{
		Driver:   "postgres",
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "postgres",
		DBName:   "unidb",
		SSLMode:  "disable",
	}))

	var accounts []Accounts
	query := "SELECT user_id, username, password, email, created_on, last_login FROM accounts"

	err := unidb.Find(&accounts, query)
	if err != nil {
		panic(err)
	}

	fmt.Println(accounts)

}
