package structutils

import (
	"reflect"
)

func GetTags(tag string, model any) []string {
	t := reflect.TypeOf(model).Elem()

	var tags []string

	for i := 0; i < t.NumField(); i++ {

		fieldTag := t.Field(i).Tag.Get(tag)
		if tag != "" {
			tags = append(tags, fieldTag)
		}
	}

	return tags
}
