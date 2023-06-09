package db

import (
    "context"

    "github.com/jackc/pgx/v5"
)

type Database struct{
    Conn *pgx.Conn
}


func Connect(connstring string) (Database, error) {
    var db Database

    ctx := context.Background()
    connConfig, err := pgx.ParseConfig(connstring)
    if err != nil {
        return db, err
    }

    db.Conn, err = pgx.ConnectConfig(ctx, connConfig)
    if err != nil {
        return db, err
    }

    return db, nil
}
