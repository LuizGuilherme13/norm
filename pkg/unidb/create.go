package unidb

import (
	"errors"
	"reflect"
	"strings"

	"github.com/LuizGuilherme13/unidb/pkg/unidb/models"
)

func (udb *UniDB) Create(dest interface{}) error {

	destValue := reflect.ValueOf(dest)
	if destValue.Kind() != reflect.Ptr || destValue.Elem().Kind() != reflect.Struct {
		return errors.New("dest must be a pointer to a struct")
	}

	destType := destValue.Elem().Type()

	var columns string
	var values []interface{}
	var qb models.QueryBuilder

	for i := 0; i < destType.NumField(); i++ {
		field := destType.Field(i)
		tag := field.Tag.Get("db")
		fieldValue := destValue.Elem().Field(i)

		if !fieldValue.IsZero() {
			columns += tag + ", "
			values = append(values, fieldValue.Interface())
		}

	}
	columns = strings.TrimSuffix(columns, ", ")

	qb.Table("accounts").Insert(columns, values...)

	err := udb.Open()
	if err != nil {
		panic(err)
	}
	defer udb.DB.Close()

	if _, err := udb.DB.Exec(qb.Q, qb.A...); err != nil {
		return err
	}

	return nil
}
