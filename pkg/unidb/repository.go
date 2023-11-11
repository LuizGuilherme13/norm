package unidb

import "github.com/jmoiron/sqlx"

// Manager is an interface what can be used to manage database
type Manager interface {
	// Find is a method what can be used to find any data from database
	Find(dest interface{}, query string, args ...interface{}) error
	// Create is a method what can be used to create any data to database
	Create()
	// OpenDB is a method what can be used to open database connection
	OpenDB() (*sqlx.DB, error)
}
