package unodatabase

import (
	"strconv"
	"strings"
)

func (udb *UniDB) Where(cond string, args ...any) *UniDB {
	if len(args) > 0 {
		for i := range args {
			pos := strconv.Itoa(len(udb.Query.A) + 1)
			cond = strings.Replace(cond, "?", "$"+pos, 1)
			udb.Query.A = append(udb.Query.A, args[i])
		}
		udb.Query.Q += " WHERE " + cond
	}
	return udb
}
