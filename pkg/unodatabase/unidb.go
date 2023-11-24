package unodatabase

import (
	"database/sql"

	"github.com/LuizGuilherme13/unodatabase/pkg/unodatabase/models"
)

type UniDB struct {
	Conn  models.DBConn
	DB    *sql.DB
	Query models.QueryBuilder
}

func New(conn models.DBConn) IUniDB {
	return &UniDB{
		Conn: conn,
	}
}

// func (udb *UniDB) formatNullColumn(col string) string {
// 	modelValue := reflect.ValueOf(udb.Query.M).Elem()

// 	fieldValue := modelValue.FieldByName(col)

// 	switch fieldValue.Type() {
// 	case reflect.TypeOf(sql.NullString{}):
// 		return fmt.Sprintf("IFNULL(%s, ''),", col)
// 	case reflect.TypeOf(sql.NullInt64{}):
// 		return fmt.Sprintf("IFNULL(%s, 0),", col)
// 	case reflect.TypeOf(sql.NullFloat64{}):
// 		return fmt.Sprintf("IFNULL(%s, 0.0),", col)
// 	case reflect.TypeOf(sql.NullBool{}):
// 		return fmt.Sprintf("IFNULL(%s, false),", col)
// 	case reflect.TypeOf(sql.NullTime{}):
// 		return fmt.Sprintf("IFNULL(%s, '0001-01-01'),", col)
// 	default:
// 		return col + ","
// 	}

// }
