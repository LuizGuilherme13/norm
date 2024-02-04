package unodatabase

import (
	"fmt"

	"github.com/LuizGuilherme13/unodatabase/pkg/unodatabase/models"
)

type Option func(*Service)

func (s *Service) WithConditions(conditions ...Option) *Service {

	for _, condition := range conditions {
		condition(s)
	}

	return s
}

func And(condition models.Condition) Option {
	return func(s *Service) {
		s.Query.Conditions = append(s.Query.Conditions,
			models.Condition{
				Condition: fmt.Sprintf("AND(%s)", condition.Condition),
				Value:     condition.Value,
			},
		)
	}
}

func Or(condition models.Condition) Option {
	return func(s *Service) {
		s.Query.Conditions = append(s.Query.Conditions,
			models.Condition{
				Condition: fmt.Sprintf("OR(%s)", condition.Condition),
				Value:     condition.Value,
			},
		)
	}
}
