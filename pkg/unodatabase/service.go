package unodatabase

import (
	"fmt"
	"reflect"

	"github.com/LuizGuilherme13/unodatabase/pkg/internal/structutils"
	"github.com/LuizGuilherme13/unodatabase/pkg/unodatabase/models"
)

type Service struct {
	r      Repository
	Query  models.Query
	Errors models.ErrorGroup
}

// NewService is a function that creates a new repository instance
func NewService(r Repository) UseCase {
	return &Service{
		r: r,
	}
}

func (s *Service) InTable(name string) *Service {
	if name == "" {
		// return nil, fmt.Errorf("the name of the table must be informed")
		s.Errors.Add(
			"unodatabase.ToModel",
			"the name of the table must be informed",
		)
	}

	s.Query.Table = name
	return s
}

func (s *Service) ToModel(model any) *Service {

	t := reflect.TypeOf(model)
	if t.Kind() != reflect.Pointer && t.Elem().Kind() != reflect.Struct {
		s.Errors.Add(
			"unodatabase.ToModel",
			fmt.Sprintf("the '%s' struct should be a pointer", t.Name()),
		)
	}

	s.Query.Dest.Self = model

	fields := structutils.GetTags("db", model)
	s.Query.Dest.Fields = append(s.Query.Dest.Fields, fields...)

	return s
}

func (s *Service) FromModel(model any) *Service {
	t := reflect.TypeOf(model)
	if t.Kind() != reflect.Pointer && t.Elem().Kind() != reflect.Struct {
		s.Errors.Add(
			"unodatabase.ToModel",
			fmt.Sprintf("the '%s' struct should be a pointer", t.Name()),
		)
	}

	s.Query.Dest.Self = model

	fields := structutils.GetTags("db", model)
	s.Query.Dest.Fields = append(s.Query.Dest.Fields, fields...)

	values := structutils.GetValues("db", model)
	for _, field := range s.Query.Dest.Fields {
		s.Query.Dest.Values = append(s.Query.Dest.Values, values[field])
	}
	return s
}

func (s *Service) Only(cols ...string) *Service {
	if len(cols) == 1 && cols[0] == "*" {
		s.Query.Fields = append(s.Query.Fields, s.Query.Dest.Fields...)
	} else if len(cols) > 0 {
		for _, col := range cols {
			for _, field := range s.Query.Dest.Fields {
				if col == field {
					s.Query.Fields = append(s.Query.Fields, field)
				}
			}
		}
	}

	return s
}

func (s *Service) Omit(cols ...string) *Service {
	for _, col := range cols {
		for _, field := range s.Query.Dest.Fields {
			if col != field {
				s.Query.Fields = append(s.Query.Fields, field)
			}
		}
	}
	return s
}

func (s *Service) Where(conditions ...models.Condition) *Service {
	s.Query.Conditions = conditions
	return s
}

func (s *Service) Find() error {

	if len(s.Errors.Errors) > 0 {
		return fmt.Errorf(s.Errors.Errors[0].String())
	}

	err := s.r.Find(s.Query)
	if err != nil {
		return fmt.Errorf("unodatabase.Find: %w", err)

	}

	return nil
}

func (s *Service) Create() {
	s.r.Create(s.Query)
}
