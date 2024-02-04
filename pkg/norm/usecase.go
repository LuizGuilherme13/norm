package norm

type UseCase interface {
	InTable(name string) *Service
	ToModel(model any) *Service
	FromModel(model any) *Service
	Only(cols ...string) *Service
	Omit(cols ...string) *Service
	WithConditions(conditions ...Option) *Service
	Find() error
	FindAll() error
	Create() error
}
