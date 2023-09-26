// Agent sends metrics as json batches to the server. Server parameters are provided by cmd line parameters
// or environment variables
package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	backoff "github.com/cenkalti/backoff/v4"
	"github.com/shipherman/go-metrics/internal/storage"
)

var sigint = make(chan os.Signal, 1)

func main() {
	// Parse cli options
	cfg, err := parseOptions()
	if err != nil {
		panic(err)
	}

	// Initiate new storage
	m := storage.New()

	// Init channels
	metricsCh := make(chan storage.MemStorage, cfg.RateLimit)
	defer close(metricsCh)

	ctx, cancel := context.WithCancel(context.Background())

	// Add wait group with RateLimit counter
	var wg sync.WaitGroup
	wg.Add(cfg.RateLimit)

	// init Channel to shutdown goroutines
	shtCh := make(chan bool)
	shtTimerCh := make(chan bool)

	// Collect data from MemStats and send to the server
	// Gather facts
	go func(timer time.Duration) {
		for {
			select {
			case <-shtTimerCh:
				log.Println("Closing timer goroutine")
				return
			default:
				time.Sleep(timer)
				readMemStats(&m, metricsCh)
			}
		}
	}(time.Second * time.Duration(cfg.PollInterval))

	// Create backoff for retrier
	b := backoff.NewExponentialBackOff()
	b.InitialInterval = time.Duration(cfg.PollInterval) * time.Second
	b.MaxElapsedTime = b.InitialInterval * time.Duration(cfg.MaxRetryInterval)

	// func for retrier
	fn := func() error {
		return ProcessBatch(context.Background(), cfg, metricsCh)
	}

	// Send metrics to the server
	for w := 1; w <= cfg.RateLimit; w++ {
		go func() {
			for {
				err := backoff.Retry(fn, b)
				if err != nil {
					log.Println(err)
				}
				select {
				case <-shtCh:
					wg.Done()
					log.Println("closing worker")
					return
				default:
					continue
				}

			}
		}()
	}

	// Gracefull shutdown here
	go func(ctx context.Context) {
		signal.Notify(sigint, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
		<-sigint
		// send true to shutdown channel to close goroutines
		for w := 1; w <= cfg.RateLimit; w++ {
			shtCh <- true
		}

		// send true to Timer goroutine
		shtTimerCh <- true

		wg.Wait()
		cancel()
		log.Println("Agent shutted down")
	}(ctx)

	<-ctx.Done()
}
