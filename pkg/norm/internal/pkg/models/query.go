package models

type Query struct {
	Model      Model
	Table      string
	Fields     []string
	Args       []any
	Conditions []Condition
}
