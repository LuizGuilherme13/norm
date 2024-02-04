package models

import "github.com/LuizGuilherme13/norm/pkg/db"

type Query struct {
	Model      Model
	Table      string
	Fields     []string
	Args       []any
	Conditions []db.Condition
}
