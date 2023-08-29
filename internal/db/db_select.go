package db

import (
	"context"
)

// Select all metrics from database
func (db *Database) SelectAll() error {
	rows, err := db.Conn.Query(context.Background(),
		`SELECT * FROM counter_metrics
        UNION
        SELECT * FROM gauge_metrics`)
	if err != nil {
		return err
	}

	for rows.Next() {
		_, err := rows.Values()
		if err != nil {
			return err
		}
	}

	return nil
}
