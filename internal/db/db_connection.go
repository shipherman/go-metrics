package db

import (
    "context"
    "fmt"

    "github.com/jackc/pgx/v5"
)

const (
    host     = "127.0.0.1"
    port     = 5432
    user     = "postgres"
    password = "pass"
    dbname   = "postgres"
)

func Connect(connstring string) (*pgx.Conn, error) {
    var conn *pgx.Conn

    ctx := context.Background()
    connConfig, err := pgx.ParseConfig(connstring)
    if err != nil {
        return conn, err
    }

    conn, err = pgx.ConnectConfig(ctx, connConfig)
    if err != nil {
        return conn, err
    }

    return conn, nil
}
