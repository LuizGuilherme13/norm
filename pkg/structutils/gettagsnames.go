package structutils

import "reflect"

func GetTagsNames(tag string, from interface{}) []string {
	fromType := reflect.TypeOf(from)

	if fromType.Kind() == reflect.Ptr {
		fromType = fromType.Elem()
	}

	if fromType.Kind() != reflect.Struct {
		return nil
	}

	var tags []string

	for i := 0; i < fromType.NumField(); i++ {
		field := fromType.Field(i)
		tags = append(tags, field.Tag.Get(tag))
	}

	return tags
}
