package structutils

import "reflect"

func GetTagsNames(tag string, from interface{}, isZero bool) ([]string, []any) {
	fromValue := reflect.ValueOf(from)
	fromType := fromValue.Elem().Type()

	var (
		tags   []string
		values []any
	)

	for i := 0; i < fromType.NumField(); i++ {
		field := fromType.Field(i)
		tag := field.Tag.Get(tag)
		fieldValue := fromValue.Elem().Field(i)

		if !fieldValue.IsZero() {
			tags = append(tags, tag)
			values = append(values, fieldValue.Interface())
		}

	}

	return tags, values
}
