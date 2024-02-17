package norm

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
	SSLMode  string
}

// Init initializes the database connection with the given settings
func (c *Config) Init() (*sql.DB, error) {
	db, err := sql.Open("postgres", c.Get())
	if err != nil {
		return nil, fmt.Errorf("postgres.Find(sql.Open): %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("postgres.Find(db.Ping): %w", err)
	}

	return db, nil
}

// Get returns the database configuration in a ready-to-use formatted string
func (c *Config) Get() string {
	return ("host=" + c.Host +
		" port=" + c.Port +
		" user=" + c.Username +
		" password=" + c.Password +
		" dbname=" + c.Database +
		" sslmode=" + c.SSLMode)
}

type Option func(*Config)

func WithHost(host string) Option {
	return func(c *Config) {
		c.Host = host
	}
}
func WithPort(port string) Option {
	return func(c *Config) {
		c.Port = port
	}
}
func WithUsername(username string) Option {
	return func(c *Config) {
		c.Username = username
	}
}
func WithPassword(password string) Option {
	return func(c *Config) {
		c.Password = password
	}
}
func WithDatabase(database string) Option {
	return func(c *Config) {
		c.Database = database
	}
}
func WithSSLMode(sslmode string) Option {
	return func(c *Config) {
		c.SSLMode = sslmode
	}
}
