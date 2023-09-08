package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	pgx "github.com/jackc/pgx/v5"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/shipherman/go-metrics/internal/storage"
	"github.com/stretchr/testify/require"

	_ "github.com/lib/pq"
)

var dbc Database
var dsn string

var db *sql.DB

func TestMain(m *testing.M) {
	// Start a new docker pool
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not construct pool: %s", err)
	}

	// Uses pool to try to connect to Docker
	err = pool.Client.Ping()
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	pg, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "15",
		Env: []string{
			"POSTGRES_PASSWORD=pass",
			"POSTGRES_USER=postgres",
			"POSTGRES_DB=postgres",
			"listen_addresses = 'localhost'",
			"listen_port = '5432'",
		},
		ExposedPorts: []string{"5432/tcp"},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"5432/tcp": {{HostIP: "localhost", HostPort: "5432/tcp"}},
		},
	}, func(config *docker.HostConfig) {
		// set AutoRemove to true so that stopped container goes away by itself
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{
			Name: "no",
		}
	})

	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	pg.Expire(120)

	hostAndPort := pg.GetHostPort("5432/tcp")
	databaseUrl := fmt.Sprintf("postgres://postgres:pass@%s/postgres?sslmode=disable", hostAndPort)

	log.Println("Connecting to database on url: ", databaseUrl)
	pool.MaxWait = 20 * time.Second
	if err = pool.Retry(func() error {
		db, err = sql.Open("postgres", databaseUrl)
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

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

	dbc.CreateTables()

	code := m.Run()

	os.Exit(code)

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

func TestWrite(t *testing.T) {
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
	t.Run("Write to DB", func(t *testing.T) {
		err := dbc.Write(s)
		require.NoError(t, err)
	})

}

func TestSelectAll(t *testing.T) {
	err := dbc.SelectAll()
	require.NoError(t, err)
}
