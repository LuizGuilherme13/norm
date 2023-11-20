package unidb

type IUniDB interface {
	Open() error
	Migrate(model interface{}) error
	Model(model any) *UniDB
	Create() error
	Get(cols ...string) error
	Where(cond string, args ...any) *UniDB
}
