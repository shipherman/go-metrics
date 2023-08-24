package db

import (
	"context"
	"fmt"
	"os"
	"testing"

	pgx "github.com/jackc/pgx/v5"
	"github.com/shipherman/go-metrics/internal/storage"
	"github.com/stretchr/testify/require"
)

var dbc Database
var dsn string

func TestMain(m *testing.M) {
	dsn = os.Getenv("DATABASE_DSN")
	ConnString, err := pgx.ParseConfig(dsn)
	// For local tests
	// "host=127.0.0.1 port=5432 dbname=postgres user=postgres password=pass"
	if err != nil {
		fmt.Printf("error during parsing conn string: %s", err.Error())
	}
	dbc.Conn, err = pgx.ConnectConfig(context.Background(), ConnString)
	if err != nil {
		fmt.Printf("error during connection to db: %s", err.Error())
	}
	dbc.createTables()
	os.Exit(m.Run())
}

func BenchmarkWrite(b *testing.B) {
	s := storage.MemStorage{
		CounterData: map[string]storage.Counter{
			"c01": 1,
			"c02": 2,
		},
		GaugeData: map[string]storage.Gauge{
			"g01": 1.01,
			"g02": 2.02,
		},
	}
	b.Run("Write to DB", func(b *testing.B) {
		err := dbc.Write(s)
		require.NoError(b, err)
	})

}

func BenchmarkSelectAll(b *testing.B) {
	err := dbc.SelectAll()
	require.NoError(b, err)
}
