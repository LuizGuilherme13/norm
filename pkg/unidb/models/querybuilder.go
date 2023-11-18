package models

import (
	"strconv"
	"strings"
)

type QueryBuilder struct {
	T string
	Q string
	A []any
}

func (qb *QueryBuilder) Table(t string) *QueryBuilder {
	qb.T = t
	return qb
}

func (qb *QueryBuilder) Get(cols ...string) *QueryBuilder {
	qb.Q = "SELECT " + strings.Join(cols, ", ") + " FROM " + qb.T

	return qb
}

func (qb *QueryBuilder) Insert(cols string, values ...any) *QueryBuilder {
	qb.Q = "INSERT INTO " + qb.T + " (" + cols + ") "
	qb.Q += "VALUES ("
	for i := 0; i < len(values); i++ {
		pos := strconv.Itoa(len(qb.A) + 1)
		qb.Q += "$" + pos + ", "
		qb.A = append(qb.A, values[i])
	}
	qb.Q = strings.TrimSuffix(qb.Q, ", ") + ")"

	return qb
}

func (qb *QueryBuilder) Update(cols map[string]any) *QueryBuilder {

	qb.Q = "UPDATE " + qb.T + " SET "
	for k, v := range cols {
		qb.Q += k + " = :" + k + ", "
		qb.A = append(qb.A, v)
	}
	qb.Q += strings.TrimSuffix(qb.Q, ", ")

	return qb

}

func (qb *QueryBuilder) Where(cond string, args ...any) *QueryBuilder {
	if len(args) > 0 {
		for i := range args {
			pos := strconv.Itoa(len(qb.A) + 1)
			cond = strings.Replace(cond, "?", "$"+pos, 1)
			qb.A = append(qb.A, args[i])
		}
		qb.Q += " WHERE " + cond
	}
	return qb
}

func (qb *QueryBuilder) And(cond string, args ...any) *QueryBuilder {
	if len(qb.A) > 0 {
		for i := range args {
			pos := strconv.Itoa(len(qb.A) + 1)
			cond = strings.Replace(cond, "?", "$"+pos, 1)
			qb.A = append(qb.A, args[i])
		}
	}
	qb.Q += " AND " + cond
	return qb
}

func (qb *QueryBuilder) Or(cond string, args ...any) *QueryBuilder {
	if len(qb.A) > 0 {
		for i := range args {
			pos := strconv.Itoa(len(qb.A) + 1)
			cond = strings.Replace(cond, "?", "$"+pos, 1)
			qb.A = append(qb.A, args[i])
		}
	}
	qb.Q += " OR " + cond
	return qb
}

func (qb *QueryBuilder) OrderBy(cols ...string) *QueryBuilder {
	qb.Q += " ORDER BY " + strings.Join(cols, ", ")
	return qb
}
