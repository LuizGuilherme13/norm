package unodatabase

import "github.com/LuizGuilherme13/unodatabase/pkg/unodatabase/models"

func (s *Service) Where(conditions ...models.Condition) *Service {
	s.Query.Conditions = conditions
	return s
}
