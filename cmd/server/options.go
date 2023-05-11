package main

import (
    "os"
    "flag"
)

type options struct {
    address string
}


func parseOptions() (options) {
    var cfg options

    flag.StringVar(&cfg.address,
                   "a", "localhost:8080",
                    "Add addres and port in format <address>:<port>")
    flag.Parse()

    // get env vars
    if a := os.Getenv("ADDRESS"); a != "" {
        cfg.address = a
    }

    return cfg
}
