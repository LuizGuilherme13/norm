package postgres

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/LuizGuilherme13/norm/pkg/norm/internal/pkg/models"
)

func (r *Repository) Find(query models.Query) error {

	db, err := r.Conn.Init()
	if err != nil {
		return fmt.Errorf("postgres.Find(sql.Open): %w", err)
	}
	defer db.Close()

	q := "SELECT " + strings.Join(query.Fields, ", ")
	q += " FROM " + query.Table
	q += " WHERE 1 = 1 "
	args := []any{}

	if len(query.Conditions) > 0 {
		for _, condition := range query.Conditions {
			q += condition.Condition
			args = append(args, condition.Value...)
		}
	}

	modelType := reflect.TypeOf(query.Model.Self).Elem()
	modelValue := reflect.ValueOf(query.Model.Self).Elem()
	numField := modelType.NumField()

	destFields := make([]any, numField)

	for i := 0; i < numField; i++ {
		fieldValue := modelValue.Field(i)

		destFields[i] = fieldValue.Addr().Interface()
	}

	err = db.QueryRow(q, args...).Scan(destFields...)
	if err != nil {
		return fmt.Errorf("postgres.Find(db.QueryRow): %w", err)
	}

	return nil
}
