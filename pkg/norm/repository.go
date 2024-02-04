package norm

import "github.com/LuizGuilherme13/norm/pkg/norm/internal/models"

type Repository interface {
	Find(query models.Query) error
	FindAll(models.Query) error
	Create(query models.Query) error
}
