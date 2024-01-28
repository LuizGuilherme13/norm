package models

type Query struct {
	Table      string
	Fields     []string
	Conditions []Condition
	Args       []any
	Dest       Dest
}
