package db

import (
	"context"
)

// Create two tables -- one for counters and one for gauge metrics
func (db *Database) CreateTables() error {
	_, err := db.Conn.Exec(context.Background(), `CREATE TABLE IF NOT EXISTS gauge_metrics(
        id serial PRIMARY KEY,
        name text,
        value double precision,
        timestamp timestamp)`)
	if err != nil {
		return err
	}

	_, err = db.Conn.Exec(context.Background(), `CREATE TABLE IF NOT EXISTS counter_metrics(
        id serial PRIMARY KEY,
        name text,
        value integer,
        timestamp timestamp)`)
	if err != nil {
		return err
	}

	return nil
}
