package structutils

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"time"
)

// ScanToModel mapeia os valores de uma linha de sql.Rows para uma struct, levando em consideração
// as tags 'db' da struct.
func ScanToModel(rows *sql.Rows, dest interface{}) error {

	// Verificar se dest é um ponteiro para uma struct
	destValue := reflect.ValueOf(dest)
	if destValue.Kind() != reflect.Ptr || destValue.Elem().Kind() != reflect.Struct {
		return errors.New("dest must be a pointer to a struct")
	}

	// Obtendo a struct de dest
	destStruct := destValue.Elem()
	destType := destValue.Elem().Type()

	// Mapeando o nome da coluna para o índice do campo correspondente na struct
	fieldIndex := make(map[string]int)
	for i := 0; i < destType.NumField(); i++ {
		field := destType.Field(i)
		columnName := field.Tag.Get("db")
		fieldIndex[columnName] = i
	}

	// Obtendo os nomes das colunas do resultado da consulta
	columns, err := rows.Columns()
	if err != nil {
		return err
	}

	// Criando uma slice de ponteiros para os valores dos campos da struct
	values := make([]interface{}, len(columns))
	for i, col := range columns {

		// Verificando se a coluna da consulta corresponde a um campo na struct
		if index, ok := fieldIndex[col]; ok {
			field := destStruct.Field(index)

			// Verificando se é possível obter o endereço do campo
			if field.CanAddr() {

				// Se o campo for do tipo time.Time, usar sql.NullTime
				if field.Type() == reflect.TypeOf(time.Time{}) {
					values[i] = new(sql.NullTime)
				} else {
					// Para outros tipos, usar o próprio endereço do campo
					values[i] = field.Addr().Interface()
				}

			} else {
				return fmt.Errorf("cannot address field '%s' in struct", col)
			}

		} else {
			return fmt.Errorf("column '%s' not found in struct", col)
		}
	}

	// Escaneando os valores de rows para o slice de ponteiros
	if err := rows.Scan(values...); err != nil {
		return err
	}

	return nil
}
