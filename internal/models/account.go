package models

import (
	"database/sql"
	"time"
)

type Account struct {
	// UserID    int64        `db:"user_id"`
	UserName  string       `db:"username"`
	Password  string       `db:"password"`
	Email     string       `db:"email"`
	CreatedOn time.Time    `db:"created_on"`
	LastLogin sql.NullTime `db:"last_login"`
}
