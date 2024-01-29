package unodatabase

import (
	"fmt"
	"reflect"

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

	typeOf := reflect.TypeOf(model)

	if typeOf.Kind() != reflect.Pointer || typeOf.Elem().Kind() != reflect.Struct {
		s.Errors.Add(
			"unodatabase.ToModel",
			fmt.Sprintf("the '%s' struct should be a pointer", typeOf.Name()),
		)
	} else {

		s.Query.Dest.Self = model

		element := typeOf.Elem()

		for i := 0; i < element.NumField(); i++ {
			s.Query.Dest.Fields = append(s.Query.Dest.Fields, models.Field{
				Column: element.Field(i).Tag.Get("db"),
			})
		}
	}

	return s
}

func (s *Service) FromModel(model any) *Service {
	typeOf := reflect.TypeOf(model)
	valueOf := reflect.ValueOf(model).Elem()

	if typeOf.Kind() != reflect.Pointer || typeOf.Elem().Kind() != reflect.Struct {
		s.Errors.Add(
			"unodatabase.FromModel",
			fmt.Sprintf("the '%s' struct should be a pointer", typeOf.Name()),
		)
	} else {

		s.Query.Dest.Self = model

		element := typeOf.Elem()

		for i := 0; i < element.NumField(); i++ {
			s.Query.Dest.Fields = append(s.Query.Dest.Fields, models.Field{
				Column: element.Field(i).Tag.Get("db"),
				Value:  valueOf.Field(i).Interface(),
			})
		}
	}

	return s
}

func (s *Service) Only(cols ...string) *Service {
	if len(cols) == 1 && cols[0] == "*" {
		for _, field := range s.Query.Dest.Fields {
			s.Query.Fields = append(s.Query.Fields, field.Column)
			s.Query.Args = append(s.Query.Args, field.Value)
		}
	} else if len(cols) > 0 {
		for _, col := range cols {
			var added bool
			for _, field := range s.Query.Dest.Fields {
				if col == field.Column {
					added = true
					s.Query.Fields = append(s.Query.Fields, field.Column)
					s.Query.Args = append(s.Query.Args, field.Value)
				}
			}
			if !added {
				typeOf := reflect.TypeOf(s.Query.Dest.Self)
				s.Errors.Add(
					"unodatabase.Only",
					fmt.Sprintf(
						"apparently field '%s' does not exist in struct '%s'", col, typeOf.Elem().Name()),
				)
			}
		}
	}

	return s
}

func (s *Service) Omit(cols ...string) *Service {
	for _, col := range cols {
		for _, field := range s.Query.Dest.Fields {
			if col != field.Column {
				s.Query.Fields = append(s.Query.Fields, field.Column)
				s.Query.Args = append(s.Query.Args, field.Value)

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

func (s *Service) Create() error {

	if len(s.Errors.Errors) > 0 {
		return fmt.Errorf(s.Errors.Errors[0].String())
	}
	err := s.r.Create(s.Query)
	if err != nil {
		return fmt.Errorf("unodatabase.Create: %w", err)

	}
	return nil
}
