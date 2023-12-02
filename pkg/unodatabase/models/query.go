package models

type Query struct {
	Table      string
	Cols       []string
	Conditions []Condition
	Args       []any
	Dest       Dest
}
