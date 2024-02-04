package models

import "fmt"

type ErrorGroup struct {
	Errors []Error
}

type Error struct {
	Place       string
	Description string
}

func (e *ErrorGroup) Add(place, desc string) {
	e.Errors = append(e.Errors, Error{Place: place, Description: desc})
}

func (e *Error) String() string {
	return fmt.Sprintf("%s:%s", e.Place, e.Description)
}
