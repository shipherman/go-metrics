// Agent sends metrics as json batches to the server. Server parameters are provided by cmd line parameters
// or environment variables
package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/shipherman/go-metrics/internal/storage"
)

type Sender func(context.Context, Options, chan storage.MemStorage) error

// Retry request on error "retries" times.
func Retry(sender Sender, retries int, delay time.Duration) Sender {
	return func(ctx context.Context, cfg Options, metricsCh chan storage.MemStorage) error {
		for r := 0; ; r++ {
			err := sender(ctx, cfg, metricsCh)
			if err == nil || r >= retries {
				// Return when there is no error or the maximum amount
				// of retries is reached.
				return err
			}

			log.Printf("Function call failed, retrying in %v", delay)

			// Increase delay
			delay = delay + time.Second*2

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

	// Initiate new storage
	m := storage.New()

	// Init channels
	done := make(chan struct{})
	metricsCh := make(chan storage.MemStorage, cfg.RateLimit)
	defer close(metricsCh)

	// Collect data from MemStats and send to the server
	// Gather facts
	go func(timer time.Duration) {
		for {
			time.Sleep(timer)
			readMemStats(&m, metricsCh)
		}
	}(time.Second * time.Duration(cfg.PollInterval))

	// Send metrics to the server
	for w := 1; w <= cfg.RateLimit; w++ {
		go func(timer time.Duration) {
			for {
				time.Sleep(timer)
				fn := Retry(ProcessBatch, 3, 1*time.Second)
				err := fn(context.Background(), cfg, metricsCh)
				if err != nil {
					log.Println(err)
				}
			}
		}(time.Second * time.Duration(cfg.ReportInterval))
	}

	// Gracefull shutdown
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		<-sigint

		close(done)
	}()

	<-done
}
