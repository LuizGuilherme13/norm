package norm

import "fmt"

func (s *Service) FindAll() error {

	if len(s.Errors.Errors) > 0 {
		return fmt.Errorf(s.Errors.Errors[0].String())
	}

	err := s.r.FindAll(s.Query)
	if err != nil {
		return fmt.Errorf("norm.Find: %w", err)

	}

	return nil
}
