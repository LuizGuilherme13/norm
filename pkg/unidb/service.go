package unidb

import "fmt"

type Service interface {
	Reader
	Writer
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) Find() {}

func (s *service) FindAll(query string, args ...any) ([]map[string]interface{}, error) {

	result, err := s.repo.FindAll(query, args...)
	if err != nil {
		return nil, fmt.Errorf("service.FindAll: %w", err)
	}

	return result, nil
}

func (s *service) New() {}

func (s *service) Edit() {}

func (s *service) Delete() {}
