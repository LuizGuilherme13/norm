package unodatabase

import "github.com/LuizGuilherme13/unodatabase/pkg/unodatabase/models"

type UseCase interface {
	InTable(name string) *Service
	ToModel(model any) *Service
	FromModel(model any) *Service
	Only(cols ...string) *Service
	Omit(cols ...string) *Service
	Where(conditions ...models.Condition) *Service
	Find() error
	Create()
}
