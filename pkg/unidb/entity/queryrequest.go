package entity

type QueryBuilder struct {
	Query string
	Args  []any
}

type QueryRequest struct {
	Table   string
	Columns []any
	QueryBuilder
}

type QueryResult struct {
	LastInsertId int64
	RowsAffected int64
	Rows         []interface{}
	Row          interface{}
}

func (q *QueryBuilder) Write(str string, args ...any) {
	q.Query += str
	if len(args) > 0 {
		q.Args = append(q.Args, args...)
	}
}
