package models

// FromModels is a struct what can be used to create a query from a model
type FromModel struct {
	Query        QueryBuilder
	Cols         []string
	Placeholders []string
}
