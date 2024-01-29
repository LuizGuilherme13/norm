package models

type Dest struct {
	Self   any
	Fields []Field
}

type Field struct {
	Column string
	Value  any
}
