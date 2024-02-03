package unodatabase

import (
	"fmt"
	"reflect"

	"github.com/LuizGuilherme13/unodatabase/pkg/unodatabase/models"
)

func (s *Service) ToModel(model any) *Service {

	typeOf := reflect.TypeOf(model)

	if typeOf.Kind() != reflect.Pointer || typeOf.Elem().Kind() != reflect.Struct {
		s.Errors.Add(
			"unodatabase.ToModel",
			fmt.Sprintf("the '%s' struct should be a pointer", typeOf.Name()),
		)
	} else {

		s.Query.Model.Self = model

		element := typeOf.Elem()

		for i := 0; i < element.NumField(); i++ {
			s.Query.Model.Fields = append(s.Query.Model.Fields, models.Field{
				Column: element.Field(i).Tag.Get("db"),
			})
		}
	}

	return s
}
