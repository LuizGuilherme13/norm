package entity

import (
	"time"
)

type User struct {
	ID        int64     `db:"user_id"`
	Name      string    `db:"username"`
	Password  string    `db:"password"`
	Email     string    `db:"email"`
	CreatedOn time.Time `db:"created_on"`
	LastLogin time.Time `db:"last_login"`
}
