// Provides commandline and environment options parsing
package options

import (
	"flag"

	"github.com/caarlos0/env"
)

type Options struct {
	Address   string `env:"ADDRESS"`
	Interval  int    `env:"STORE_INTERVAL"`
	Filename  string `env:"FILE_STORAGE_PATH"`
	Restore   bool   `env:"RESTORE"`
	DBDSN     string `env:"DATABASE_DSN"`
	Key       string `env:"KEY"`
	CryptoKey string `env:"CRYPTO_KEY"`
}

func ParseOptions() (Options, error) {
	var cfg Options

	flag.StringVar(&cfg.Address,
		"a", "localhost:8080",
		"Add address and port in format <address>:<port>")
	flag.IntVar(&cfg.Interval,
		"i", 300,
		"Saving metrics to file interval")
	flag.StringVar(&cfg.Filename,
		"f", "/tmp/metrics-db.json",
		"File path")
	flag.BoolVar(&cfg.Restore,
		"r", true,
		"Restore metrics value from file")
	flag.StringVar(&cfg.DBDSN,
		"d",
		"",
		"Connection string in Postgres format")
	flag.StringVar(&cfg.Key, "k", "", "Sing key")
	flag.StringVar(&cfg.CryptoKey, "crypto-key", "",
		"Private key path")
	flag.Parse()

	// get env vars
	err := env.Parse(&cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}
