package postgres

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/LuizGuilherme13/norm/pkg/norm/internal/pkg/models"
)

func (r *Repository) FindAll(query models.Query) error {
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
			args = append(args, condition.Value)
		}
	}

	rows, err := db.Query(q, args...)
	if err != nil {
		return fmt.Errorf("postgres.FindAll(db.Query): %w", err)
	}
	defer rows.Close()

	// Obtém o tipo subjacente da slice de structs
	sliceType := reflect.TypeOf(query.Model.Self).Elem().Elem()
	// Cria uma slice vazia com o mesmo tipo da slice passada
	slice := reflect.MakeSlice(reflect.SliceOf(sliceType), 0, 0)

	// Obtém o valor da slice vazia
	sliceValue := reflect.ValueOf(slice.Interface())

	for rows.Next() {
		// Cria uma nova struct do tipo subjacente da slice
		modelValue := reflect.New(sliceType).Elem()
		numField := modelValue.NumField()
		destFields := make([]any, numField)

		for i := 0; i < numField; i++ {
			fieldValue := modelValue.Field(i)
			destFields[i] = fieldValue.Addr().Interface()
		}

		if err := rows.Scan(destFields...); err != nil {
			return fmt.Errorf("postgres.FindAll(rows.Scan): %w", err)
		}

		// Adiciona a struct ao slice
		sliceValue = reflect.Append(sliceValue, modelValue)
	}

	// Atualiza a slice passada com os valores adicionados
	reflect.ValueOf(query.Model.Self).Elem().Set(sliceValue)

	if err := rows.Err(); err != nil {
		return fmt.Errorf("postgres.FindAll(rows.Err): %w", err)
	}

	return nil
}
