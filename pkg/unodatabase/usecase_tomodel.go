package unodatabase

import (
	"fmt"
	"reflect"

	"github.com/LuizGuilherme13/unodatabase/pkg/unodatabase/models"
)

func (s *Service) ToModel(model any) *Service {

	typeOf := reflect.TypeOf(model)
	elem := typeOf.Elem()

	if typeOf.Kind() != reflect.Pointer {
		s.Errors.Add(
			"unodatabase.ToModel",
			fmt.Sprintf("the '%s' should be a pointer", typeOf.Name()),
		)

		if elem.Kind() != reflect.Slice && elem.Elem().Kind() != reflect.Struct {
			s.Errors.Add(
				"unodatabase.ToModel",
				"the param 'model' should be a struct",
			)
		}
		if elem.Kind() == reflect.Slice && elem.Elem().Kind() != reflect.Struct {
			s.Errors.Add(
				"unodatabase.ToModel",
				"the param 'model' should be a slice of struct",
			)
		}

	} else {

		s.Query.Model.Self = model

		var element reflect.Type

		if elem.Kind() != reflect.Slice {
			element = elem
		} else {
			element = elem.Elem()
		}

		for i := 0; i < element.NumField(); i++ {
			s.Query.Model.Fields = append(s.Query.Model.Fields, models.Field{
				Column: element.Field(i).Tag.Get("db"),
			})
		}
	}

	return s
}
