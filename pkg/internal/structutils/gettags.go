package structutils

import (
	"reflect"
)

func GetTags(tag string, model any) []string {
	modelValue := reflect.ValueOf(model).Elem()
	modelType := modelValue.Type()

	var tags []string

	for i := 0; i < modelType.NumField(); i++ {
		field := modelType.Field(i)

		t := field.Tag.Get(tag)

		if t != "" {
			tags = append(tags, t)
		}
	}

	return tags
}
