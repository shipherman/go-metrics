// Provides commandline and environment options parsing
package options

import (
	"encoding/json"
	"flag"
	"net"
	"os"

	"github.com/caarlos0/env"
)

type Options struct {
	Address             string `env:"ADDRESS" json:"address"`
	Interval            int    `env:"STORE_INTERVAL" json:"store_interval"`
	Filename            string `env:"FILE_STORAGE_PATH" json:"store_file"`
	Restore             bool   `env:"RESTORE" json:"restore"`
	DBDSN               string `env:"DATABASE_DSN" json:"database_dsn"`
	Key                 string `env:"KEY" json:"key"`
	CryptoKey           string `env:"CRYPTO_KEY" json:"crypto_key"`
	ConfigPath          string `env:"CONFIG"`
	TrustedSubnet       string `env:"TRUSTED_SUBNET" json:"trusted_subnet"`
	TrustedSubnetParsed *net.IPNet
}

func ParseOptions() (Options, error) {
	var cfg Options

	flag.StringVar(&cfg.ConfigPath,
		"c", "",
		"Configuration file path")
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
	flag.StringVar(&cfg.TrustedSubnet,
		"t",
		"192.168.1.0/24",
		"Trusted subnet defines allowed subnet to receive requests")
	flag.Parse()

	// get env vars
	err := env.Parse(&cfg)
	if err != nil {
		return cfg, err
	}

	if cfg.ConfigPath != "" {
		cfg, err = ReadConfigFile(cfg.ConfigPath)
		if err != nil {
			return cfg, err
		}
	}

	// Executing subnet from TrustedSubnet parameter
	_, cfg.TrustedSubnetParsed, err = net.ParseCIDR(cfg.TrustedSubnet)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}

func ReadConfigFile(path string) (cfg Options, err error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return cfg, err
	}

	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}
