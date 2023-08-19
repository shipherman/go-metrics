package db

import (
	"context"
)

func (db *Database) createTables() error {
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
