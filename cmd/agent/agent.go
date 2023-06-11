package main

import (
    "log"
    "time"
    "context"
    "github.com/shipherman/go-metrics/internal/storage"

)

type Sender func(context.Context, string, storage.MemStorage) error


func Retry(sender Sender, retries int, delay time.Duration) Sender {
    return func(ctx context.Context, serverAddress string, m storage.MemStorage) error {
        for r := 0; ; r++ {
            err := sender(ctx, serverAddress, m)
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
    //parse cli options
    cfg, err := parseOptions()
    if err != nil {
        panic(err)
    }

    // initiate tickers
    pollTicker := time.NewTicker(time.Second * time.Duration(cfg.PollInterval))
	defer pollTicker.Stop()
    reportTicker := time.NewTicker(time.Second * time.Duration(cfg.ReportInterval))
	defer reportTicker.Stop()

    //initiate new storage
    m := storage.New()

    //collect data from MemStats and send to the server
    for {
        select {
        case <-pollTicker.C:
            readMemStats(&m)
        case <-reportTicker.C:
            fn := Retry(ProcessBatch, 3, 1*time.Second)
            err := fn(context.Background(), cfg.ServerAddress, m)
            if err != nil {
                log.Println(err)
            }
        }
    }
}
