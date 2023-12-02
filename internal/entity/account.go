package entity

import (
	"time"
)

type Account struct {
	UserID    int64     `db:"user_id"`
	UserName  string    `db:"username"`
	Password  string    `db:"password"`
	Email     string    `db:"email"`
	CreatedOn time.Time `db:"created_on"`
	LastLogin time.Time `db:"last_login"`
}
