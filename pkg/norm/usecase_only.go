package norm

import "fmt"

func (s *Service) Only(cols ...string) *Service {
	if len(cols) == 1 && cols[0] == "*" {
		for _, field := range s.Query.Model.Fields {
			s.Query.Fields = append(s.Query.Fields, field.Column)
			s.Query.Args = append(s.Query.Args, field.Value)
		}
	} else if len(cols) > 0 {
		for _, col := range cols {
			var added bool
			for _, field := range s.Query.Model.Fields {
				if col == field.Column {
					added = true
					s.Query.Fields = append(s.Query.Fields, field.Column)
					s.Query.Args = append(s.Query.Args, field.Value)
				}
			}
			if !added {
				s.Errors.Add(
					"norm.Only",
					fmt.Sprintf(
						"apparently field '%s' does not exist in struct '%s'", col, s.Query.Model.Name()),
				)
			}
		}
	}

	return s
}
