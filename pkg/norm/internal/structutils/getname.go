package structutils

import (
	"reflect"
)

// GetName retorna o nome da struct
func GetName(model interface{}) string {
	modelType := reflect.TypeOf(model)

	if modelType.Kind() == reflect.Ptr {
		modelType = modelType.Elem()
	}

	return modelType.Name()

}
