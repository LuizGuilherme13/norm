package unodatabase

import "github.com/LuizGuilherme13/unodatabase/pkg/unodatabase/models"

type Repository interface {
	Find(query models.Query) error
	FindAll(models.Query) error
	Create(query models.Query) error
}
