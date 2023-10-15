// Agent options, such as server addres, requests intervls, etc.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/caarlos0/env"
)

type Options struct {
	ServerAddress    string `env:"ADDRESS" json:"server_address"`
	PollInterval     int    `env:"POLL_INTERVAL" json:"poll_interval"`
	MaxRetryInterval int    `env:"MAX_RETRY_INTERVAL" json:"max_retry_interval"`
	ReportInterval   int    `env:"REPORT_INTERVAL" json:"report_interval"`
	RateLimit        int    `env:"RATE_LIMIT" json:"rate_limit"`
	Key              string `env:"KEY" json:"key"`
	CryptoKey        string `env:"CRYPTO_KEY" json:"crypto_key"`
	KeyByte          []byte
	ConfigPath       string `env:"CONFIG"`
	Encrypt          bool
	RealIP           string
}

func ParseOptions() (Options, error) {
	var cfg Options
	cfg.Encrypt = false

	// Parse environment variables
	err := env.Parse(&cfg)
	if err != nil {
		return cfg, err
	}

	// Parse cli parameters
	flag.Visit(func(f *flag.Flag) {
		if f.Name == "config" {
			flag.StringVar(&cfg.ConfigPath,
				"config", "",
				"Configuration file path")
		}
	})
	flag.IntVar(&cfg.PollInterval, "p", 2,
		"Frequensy in seconds for collecting metrics")
	flag.IntVar(&cfg.MaxRetryInterval, "m", 4,
		"Max interval to wait aswer from server")
	flag.IntVar(&cfg.ReportInterval, "r", 10,
		"Frequensy in seconds for sending report to the server")
	flag.StringVar(&cfg.ServerAddress, "a", "localhost:8080",
		"Address of the server to send metrics")
	flag.StringVar(&cfg.Key, "k", "",
		"Encryption key")
	flag.StringVar(&cfg.CryptoKey, "crypto-key", "",
		"Public key path")
	flag.IntVar(&cfg.RateLimit, "l", 3,
		"Rate Limit")

	// Read configuration file
	if cfg.ConfigPath != "" {
		cfg, err = readConfigFile(cfg.ConfigPath)
		if err != nil {
			return cfg, err
		}
	}

	// Save endpoint address
	cfg.ServerAddress = strings.Join([]string{"http:/", cfg.ServerAddress, "updates/"}, "/")

	if cfg.Key != "" {
		cfg.Encrypt = true
		cfg.KeyByte = []byte(cfg.Key)
	}

	// Get Agent IP address
	ip, err := getIPAddr()
	if err != nil {
		return cfg, err
	}
	cfg.RealIP = ip

	// View config in console
	fmt.Println(cfg)

	return cfg, nil
}

// Gets first non local loop Network interface address
func getIPAddr() (string, error) {
	addresses, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, addr := range addresses {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.String(), nil
			}
		}
	}
	return "", fmt.Errorf("no ethernet adapter found")
}

// Reads configuration file
func readConfigFile(path string) (cfg Options, err error) {
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
