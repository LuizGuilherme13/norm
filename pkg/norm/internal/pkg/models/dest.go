package models

import "reflect"

type Model struct {
	Self   any
	Fields []Field
}

func (m *Model) Name() string {
	typeOf := reflect.TypeOf(m.Self)

	if typeOf.Kind() == reflect.Ptr {
		typeOf = typeOf.Elem()
	}

	return typeOf.Name()
}

type Field struct {
	Column string
	Value  any
}
