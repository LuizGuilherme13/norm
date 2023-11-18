package structutils

import (
	"database/sql"
	"errors"
	"reflect"
)

func Scan(rows *sql.Rows, dest interface{}) error {

	destValue := reflect.ValueOf(dest)
	if destValue.Kind() != reflect.Ptr || destValue.Elem().Kind() != reflect.Struct {
		return errors.New("dest must be a pointer to a struct")
	}

	destType := destValue.Elem().Type()

	columns, err := rows.Columns()
	if err != nil {
		return err
	}

	fieldIndex := make(map[string]int)

	for i := 0; i < destType.NumField(); i++ {
		field := destType.Field(i)
		columnName := field.Tag.Get("db")

		fieldIndex[columnName] = i
	}

	// Criar um slice de ponteiros para valores para passar para o Scan
	destSlice := make([]interface{}, len(columns))
	for i := range destSlice {
		destSlice[i] = new(interface{})
	}

	if err := rows.Scan(destSlice...); err != nil {
		return err
	}

	for i, columnName := range columns {
		index, ok := fieldIndex[columnName]
		if !ok {
			// A coluna não corresponde a nenhum campo na struct, então a ignoramos
			continue
		}

		destField := destValue.Elem().Field(index)
		value := destSlice[i].(*interface{})

		if value != nil {
			fieldType := destField.Type()
			newValue := reflect.New(fieldType).Elem()

			newValue.Set(reflect.ValueOf(*value))

			destField.Set(newValue)
		}
	}

	return nil
}
