package database

type Repository interface {
	Into(table string) Repository
	WithModel(model any) Repository
	Find(cols ...string) error
	Where(condition string, values ...any) Repository
	New() error
}
