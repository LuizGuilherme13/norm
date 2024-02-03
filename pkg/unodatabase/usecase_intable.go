package unodatabase

func (s *Service) InTable(name string) *Service {
	if name == "" {
		s.Errors.Add(
			"unodatabase.ToModel",
			"the name of the table must be informed",
		)
	}

	s.Query.Table = name
	return s
}
