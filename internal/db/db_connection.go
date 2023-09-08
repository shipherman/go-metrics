// Package inplement connection to DB and requests to it
package db

import (
	"context"

	pgx "github.com/jackc/pgx/v5"
)

type Database struct {
	Conn *pgx.Conn
}

// Connect to Postgre DB using connection string provided by cmd arguments or environment variable
func Connect(ctx context.Context, connstring string) (Database, error) {
	var db Database

	connConfig, err := pgx.ParseConfig(connstring)
	if err != nil {
		return db, err
	}

	db.Conn, err = pgx.ConnectConfig(ctx, connConfig)
	if err != nil {
		return db, err
	}

	err = db.CreateTables()
	if err != nil {
		return db, err
	}

	return db, nil
}

// Close DB connection
func (db *Database) Close() {
	db.Conn.Close(context.Background())
}
