package postgres

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/LuizGuilherme13/norm/pkg/norm/internal/pkg/models"
)

func (r *Repository) Create(query models.Query) error {
	db, err := sql.Open(r.Conn.Driver, r.Conn.String())
	if err != nil {
		return fmt.Errorf("postgres.Create(sql.Open): %w", err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("postgres.Create(db.Begin): %w", err)
	}
	defer tx.Rollback()

	q := "INSERT INTO " + query.Table
	q += " (" + strings.Join(query.Fields, ",") + ")"
	q += " VALUES("

	args := []any{}
	for i := range query.Fields {
		q += fmt.Sprintf("$%d, ", i+1)
		args = append(args, query.Args[i])
	}
	q = strings.TrimSuffix(q, ", ") + ")"

	_, err = tx.Exec(q, args...)
	if err != nil {
		return fmt.Errorf("postgres.Create(tx.Exec): %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("postgres.Create(tx.Commit): %w", err)
	}

	return nil
}
