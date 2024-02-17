package norm

import (
	"fmt"

	"github.com/LuizGuilherme13/norm/pkg/norm/internal/pkg/models"
)

type Condition func(*Service)

func (s *Service) WithConditions(conditions ...Condition) *Service {

	for _, condition := range conditions {
		condition(s)
	}

	return s
}

func And(condition string, value ...any) Condition {
	return func(s *Service) {
		c := models.Condition{Condition: fmt.Sprintf("AND(%s)", condition)}
		c.Value = append(c.Value, value...)

		s.Query.Conditions = append(s.Query.Conditions, c)
	}
}

func Or(condition string, value ...any) Condition {
	return func(s *Service) {
		s.Query.Conditions = append(s.Query.Conditions,
			models.Condition{
				Condition: fmt.Sprintf("OR(%s)", condition),
				Value:     value,
			},
		)
	}
}
