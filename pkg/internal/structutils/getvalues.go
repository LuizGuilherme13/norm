package structutils

import "reflect"

func GetValues(tag string, model any) map[string]any {
	modelValue := reflect.ValueOf(model).Elem()
	modelType := modelValue.Type()

	values := make(map[string]any)

	for i := 0; i < modelType.NumField(); i++ {
		field := modelValue.Field(i)

		t := modelType.Field(i).Tag.Get(tag)

		values[t] = field.Interface()
	}

	return values

}
