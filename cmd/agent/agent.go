package main

import (
    "log"
    "time"
    "context"
    "github.com/shipherman/go-metrics/internal/storage"

)

type Sender func(context.Context, string, chan storage.MemStorage) error

var encrypt bool
var key []byte


func Retry(sender Sender, retries int, delay time.Duration) Sender {
    return func(ctx context.Context, serverAddress string, metricsCh chan storage.MemStorage) error {
        for r := 0; ; r++ {
            err := sender(ctx, serverAddress, metricsCh)
            if err == nil || r >= retries {
                // Return when there is no error or the maximum amount
                // of retries is reached.
                return err
            }

            log.Printf("Function call failed, retrying in %v", delay)

            // Increase delay
            delay = delay + time.Second * 2

            select {
            case <-time.After(delay):
            case <-ctx.Done():
                return ctx.Err()
            }
        }
    }
}

func main() {
    // Parse cli options
    cfg, err := parseOptions()
    if err != nil {
        panic(err)
    }

    if cfg.Key != "" {
        encrypt = true
        key = []byte(cfg.Key)
    }

    // Initiate tickers
    pollTicker := time.NewTicker(time.Second * time.Duration(cfg.PollInterval))
	defer pollTicker.Stop()
    reportTicker := time.NewTicker(time.Second * time.Duration(cfg.ReportInterval))
	defer reportTicker.Stop()

    // Initiate new storage
    m := storage.New()

    // Init channels
    metricsCh := make(chan storage.MemStorage, cfg.RateLimit)
    defer close(metricsCh)

    // Collect data from MemStats and send to the server
    for {
        select {
        case <-pollTicker.C:
            go readMemStats(&m, metricsCh)
        case <-reportTicker.C:
            for w := 1; w <= cfg.RateLimit; w++ {
                go func() {
                    fn := Retry(ProcessBatch, 3, 1*time.Second)
                    err := fn(context.Background(), cfg.ServerAddress, metricsCh)
                    if err != nil {
                        log.Println(err)
                    }
                }()
            }
        }
    }
}
