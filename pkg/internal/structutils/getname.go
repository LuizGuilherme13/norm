package structutils

import (
	"reflect"
	"strings"
)

// GetName retorna o nome da struct em lowercase
func GetName(model interface{}) string {
	modelType := reflect.TypeOf(model)

	if modelType.Kind() == reflect.Ptr {
		modelType = modelType.Elem()
	}

	return strings.ToLower(modelType.Name())

}
