package postgres

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/LuizGuilherme13/unodatabase/pkg/internal/structutils"
)

func (db *PsqlConn) Find(cols ...string) error {
	conn, err := sql.Open(db.Conn.Driver, db.Conn.String())
	if err != nil {
		return err
	}
	defer conn.Close()

	query := fmt.Sprintf("SELECT %s FROM %s ", strings.Join(cols, ", "), db.Query.Table)

	var args []any
	if len(db.Query.Conditions) > 0 {
		for _, c := range db.Query.Conditions {
			// fmt.Println(c.Condition, c.Value)
			query += c.Condition
			args = append(args, c.Value...)
		}
	}

	fmt.Println(query, args)
	row, err := conn.Query(query, args...)
	if err != nil {
		return err
	}
	defer row.Close()

	if row.Next() {
		err := structutils.ScanToModel(row, db.Query.Dest.Self)
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf(
			"unodatabase.Get: Register not found for: %s.\n Query: %s.\n Args: %v",
			structutils.GetName(db.Query.Dest.Self), query, args,
		)
	}

	return nil
}
