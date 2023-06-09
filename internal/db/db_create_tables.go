package db

import (
    "fmt"
    "context"

//     "github.com/jackc/pgx/v5"
)

func (db *Database) CreateTables() error {
    _, err := db.Conn.Exec(context.Background(),  `CREATE TABLE gauge_metrics(
        id serial PRIMARY KEY,
        name text,
        value double precision,
        timestamp timestamp)`)
    if err != nil {
        return err
    }

    _, err = db.Conn.Exec(context.Background(),  `CREATE TABLE counter_metrics(
        id serial PRIMARY KEY,
        name text,
        value integer,
        timestamp timestamp)`)
    if err != nil {
        return err
    }

    query := `
        SELECT
            table_schema || '.' || table_name
        FROM information_schema.tables
        WHERE
            table_schema NOT IN ('pg_catalog', 'information_schema')
    `
    rows, err := db.Conn.Query(context.Background(), query)
    if err != nil {
        return err
    }

    for rows.Next() {
        fmt.Println(rows.Values())
    }

    return nil
}
