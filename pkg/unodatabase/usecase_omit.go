package unodatabase

import "fmt"

func (s *Service) Omit(cols ...string) *Service {
	for _, col := range cols {
		var added bool

		for _, field := range s.Query.Model.Fields {
			if col != field.Column {
				s.Query.Fields = append(s.Query.Fields, field.Column)
				s.Query.Args = append(s.Query.Args, field.Value)
			}
		}

		if !added {
			s.Errors.Add(
				"unodatabase.Omit",
				fmt.Sprintf(
					"apparently field '%s' does not exist in struct '%s'",
					col, s.Query.Model.Name(),
				),
			)
		}
	}
	return s
}
