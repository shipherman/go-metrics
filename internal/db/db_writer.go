package db

import (
	"context"
	"time"

	"github.com/shipherman/go-metrics/internal/storage"
)

// Write data to DB
func (db *Database) Write(s storage.MemStorage) error {
	for k, v := range s.CounterData {
		_, err := db.Conn.Exec(context.Background(),
			`INSERT INTO counter_metrics (name, value, timestamp) VALUES ($1, $2, $3)`,
			k, v, time.Now())
		if err != nil {
			return err
		}
	}

	for k, v := range s.GaugeData {
		_, err := db.Conn.Exec(context.Background(),
			`INSERT INTO gauge_metrics (name, value, timestamp) VALUES ($1, $2, $3)`,
			k, v, time.Now())
		if err != nil {
			return err
		}
	}
	return nil
}

// Write data on timer
func (db *Database) Save(t int, s storage.MemStorage) error {
	time.Sleep(time.Second * time.Duration(t))
	return db.Write(s)
}
