package norm

import (
	"fmt"
	"reflect"

	"github.com/LuizGuilherme13/norm/pkg/norm/internal/models"
)

func (s *Service) FromModel(model any) *Service {
	typeOf := reflect.TypeOf(model)
	valueOf := reflect.ValueOf(model).Elem()

	if typeOf.Kind() != reflect.Pointer || typeOf.Elem().Kind() != reflect.Struct {
		s.Errors.Add(
			"norm.FromModel",
			fmt.Sprintf("the '%s' struct should be a pointer", typeOf.Name()),
		)
	} else {

		s.Query.Model.Self = model

		element := typeOf.Elem()

		for i := 0; i < element.NumField(); i++ {
			s.Query.Model.Fields = append(s.Query.Model.Fields, models.Field{
				Column: element.Field(i).Tag.Get("db"),
				Value:  valueOf.Field(i).Interface(),
			})
		}
	}

	return s
}
