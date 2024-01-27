package database

type Repository interface {
	// Into receives the name of the table where the query will be made
	Into(table string) Repository
	// WithModel receives a pointer to any struct
	WithModel(model any) Repository
	// Find é usada para quando a query for uma busca no banco de dados
	Find(cols ...string) error
	// Where is used to create conditionals for the query
	Where(condition string, values ...any) Repository
	// New is used when the query is an insert into the database
	New() error
}
