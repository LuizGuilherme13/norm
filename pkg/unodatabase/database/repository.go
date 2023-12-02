package database

type Repository interface {
	Migrate(model any)
	Model(model any) Repository
	Get(cols ...string) error
	Where(condition string, values ...any) Repository
	Create() error
}
