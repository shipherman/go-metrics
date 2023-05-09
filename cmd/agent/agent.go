package main

import (
    "log"
    "time"
)


//server parameters
var contentType string = "text/plain"

var logger *log.Logger


func main() {
    //parse cli options
    err := parseOptions()
    if err != nil {
        panic(err)
    }

    // initiate tickers
    pollTicker := time.NewTicker(time.Second * time.Duration(options.pollInterval))
	defer pollTicker.Stop()
    reportTicker := time.NewTicker(time.Second * time.Duration(options.reportInterval))
	defer reportTicker.Stop()

    //collect data from MemStats and send to the server
    for {
        select {
        case <-pollTicker.C:
            readMemStats(&m)
        case <-reportTicker.C:
            err := ProcessReport(&m)
            if err != nil {
                log.Println(err)
            }
        }
    }
}
