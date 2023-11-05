package postgres

import "fmt"

func (r *Repository) FindAll(query string, args ...any) ([]map[string]interface{}, error) {
	err := r.Connect()
	if err != nil {
		return nil, fmt.Errorf("db.SelectRows 1: %w", err)
	}
	defer r.DB.Close()

	rows, err := r.DB.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("db.SelectRows 2: %w", err)
	}
	defer rows.Close()

	colums, err := rows.Columns()
	if err != nil {
		return nil, fmt.Errorf("db.SelectRows 3: %w", err)
	}

	values := make([]interface{}, len(colums))
	valuesPtr := make([]interface{}, len(colums))
	for i := range colums {
		valuesPtr[i] = &values[i]
	}

	result := []map[string]interface{}{}

	for rows.Next() {
		if err := rows.Scan(valuesPtr...); err != nil {
			return nil, fmt.Errorf("db.SelectRows 4: %w", err)
		}

		row := map[string]interface{}{}
		for i, col := range colums {
			row[col] = values[i]
		}

		result = append(result, row)
	}

	return result, nil

}
