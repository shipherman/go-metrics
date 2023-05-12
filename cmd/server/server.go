package main

import (
    "net/http"
    "log"
    "github.com/shipherman/go-metrics/internal/routers"
)


func main() {
    //parse cli options
    cfg := parseOptions()

    log.Println("Starting server...")

    //run server
    log.Fatal(http.ListenAndServe(cfg.address, routers.InitRouter()))
}
