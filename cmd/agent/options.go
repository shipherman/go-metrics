package main

import (
    "flag"
    "os"
    "strconv"
)

//cli options
type Options struct {
    serverAddress string `env:"ADDRESS"`
    pollInterval int `env:"POLL_INTERVAL"`
    reportInterval int `env:"REPORT_INTERVAL"`
}

var options Options


func parseOptions () error {
// 	fmt.Println(options)
    flag.IntVar(&options.pollInterval, "p", 2,
                     "Frequensy in seconds for collecting metrics")
    flag.IntVar(&options.reportInterval, "r", 10,
                     "Frequensy in seconds for sending report to the server")
    flag.StringVar(&options.serverAddress, "a", "localhost:8080",
                "Address of the server to send metrics")
    flag.Parse()

    if l := os.Getenv("ADDRESS"); l != "" {
        options.serverAddress = l
    }
    if l := os.Getenv("POLL_INTERVAL"); l != "" {
        i, err := strconv.Atoi(l)
        if err != nil {
            return err
        }
        options.pollInterval = i
    }
    if l := os.Getenv("REPORT_INTERVAL"); l != "" {
        i, err := strconv.Atoi(l)
        if err != nil {
            return err
        }
        options.reportInterval = i
    }
    return nil
}
