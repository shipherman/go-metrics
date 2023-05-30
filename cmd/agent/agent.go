// Agent
package main

import (
    "log"
    "time"
    "github.com/shipherman/go-metrics/internal/storage"

)


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
            err := ProcessReport(cfg.ServerAddress, m)
            if err != nil {
                log.Println(err)
            }
        }
    }
}
