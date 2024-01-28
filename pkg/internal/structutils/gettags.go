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
		// value := modelValue.Field(i)

		dbTagName := field.Tag.Get(tag)

		// postTag := field.Tag.Get("post")
		// getTag := field.Tag.Get("get")

		if dbTagName != "" {

			tags = append(tags, dbTagName)
		}
	}

	return tags
}
