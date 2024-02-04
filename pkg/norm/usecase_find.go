package norm

import "fmt"

func (s *Service) Find() error {

	if len(s.Errors.Errors) > 0 {
		return fmt.Errorf(s.Errors.Errors[0].String())
	}

	err := s.r.Find(s.Query)
	if err != nil {
		return fmt.Errorf("norm.Find: %w", err)

	}

	return nil
}
