package norm

import "fmt"

func (s *Service) Create() error {

	if len(s.Errors.Errors) > 0 {
		return fmt.Errorf(s.Errors.Errors[0].String())
	}
	err := s.r.Create(s.Query)
	if err != nil {
		return fmt.Errorf("norm.Create: %w", err)

	}
	return nil
}
